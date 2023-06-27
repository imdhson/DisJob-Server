package modules

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func will_send_contains(input []Dj_jobs_detail, oid primitive.ObjectID) bool {
	for _, v := range input {
		if v.ID == oid {
			return true
		}
	}
	return false
}

func contains(input []string, v string) bool {
	if len(input) <= 1 { //빈칸일경우 무조건 true 반환
		return true
	}
	for _, v2 := range input {
		if v2 == v {
			return true
		}
	}
	return false
}
func type_union(t1 string, t2 string, t3 string) []string { //장애유형을 받아서 교집합만 배열로 반환해줌
	var rst []string
	t1 = strings.ReplaceAll(t1, " ", "") // space가 있으면 소거
	t2 = strings.ReplaceAll(t2, " ", "")
	t3 = strings.ReplaceAll(t3, " ", "")
	t1_splited := strings.Split(t1, ",") //,를 기준으로 split
	if len(t1_splited) <= 1 {            //1번 마저도 빈칸이면 모든 부위 사용 가능
		t1_splited = []string{"팔", "다리", "시각", "음성", "귀"}
	}
	t2_splited := strings.Split(t2, ",")
	t3_splited := strings.Split(t3, ",")

	for _, v := range t1_splited {
		if contains(t2_splited, v) && contains(t3_splited, v) {
			rst = append(rst, v)
		}
	}
	return rst
}
func AIListSender(w http.ResponseWriter, r *http.Request) { //메인화면 직장 리스트
	err := godotenv.Load()
	Critical(err)
	URI := os.Getenv("MONGODB_URI")
	if URI == "" {
		Critical(err)
	}
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	Critical(err)
	defer func() {
		err := db.Disconnect(context.TODO())
		Critical(err)
	}()

	//filter에 적용할 user의 데이터를 가져옴
	user_struct := OidTOuser_struct(SessionTO_oid(w, r))
	splited_loc := strings.Split(user_struct.Settings.Loc, " ")
	if len(splited_loc) <= 1 { //인덱스 런타임 에러 방지
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err_msg := map[string]string{"error": "users 관련 처리 중 오류 발생"}
		err_msg_json, _ := json.MarshalIndent(err_msg, " ", "	")
		w.Write(err_msg_json)
		return
	}

	coll_avty := db.Database("dj_jobs").Collection("type_availability")
	var typeavt [3]Dj_jobs_typeavt
	err = coll_avty.FindOne(context.TODO(), bson.D{{"종류", user_struct.Settings.Type1}}).Decode(&typeavt[0])
	ErrOK(err)
	err = coll_avty.FindOne(context.TODO(), bson.D{{"종류", user_struct.Settings.Type2}}).Decode(&typeavt[1])
	ErrOK(err)
	err = coll_avty.FindOne(context.TODO(), bson.D{{"종류", user_struct.Settings.Type3}}).Decode(&typeavt[2])
	ErrOK(err)

	avt_unioned := type_union(typeavt[0].Availability, typeavt[1].Availability, typeavt[2].Availability) //교집합 구하기
	fmt.Printf("avt0: %v\navt1: %v\navt2: %v\n", typeavt[0].Availability, typeavt[1].Availability, typeavt[2].Availability)
	//현재 개발중
	fmt.Println("avt unioned", avt_unioned)

	coll := db.Database("dj_jobs").Collection("job_list")
	//쿼리
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"사업장 주소", bson.D{{"$regex", splited_loc[0]}}}},
			bson.D{{"필수부위", bson.D{{"$regex", ""}}}},
		}}}

	cursor, err := coll.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	var will_send []Dj_jobs_detail
	App_List_num := 0
	for cursor.Next(context.TODO()) {
		var dbres Dj_jobs_detail
		cursor.Decode(&dbres)
		if !will_send_contains(will_send, dbres.ID) { //will_send에 이미 포함되지 않은 경우만 append
			dbres.App_List_num = App_List_num
			will_send = append(will_send, dbres)
			App_List_num++
		}
	}
	ErrOK(err)

	will_send_json, _ := json.MarshalIndent(will_send, " ", "	")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(will_send_json)
}
