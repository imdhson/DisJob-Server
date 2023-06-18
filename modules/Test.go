package modules

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SampleAIList(w http.ResponseWriter, r *http.Request) {
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
	coll := db.Database("dj_jobs").Collection("job_list")

	//쿼리1
	filter := bson.D{{"사업장명", "서울아산병원"}}
	var dbres_1 Dj_jobs_detail
	err = coll.FindOne(context.TODO(), filter).Decode(&dbres_1)
	ErrOK(err)
	dbres_1.App_List_num = 1
	// 쿼리2
	filter = bson.D{{"사업장명", "용인시청"}}
	var dbres_2 Dj_jobs_detail
	err = coll.FindOne(context.TODO(), filter).Decode(&dbres_2)
	ErrOK(err)
	dbres_2.App_List_num = 2
	//쿼리3
	filter = bson.D{{"사업장명", "강원랜드(주)"}}
	var dbres_3 Dj_jobs_detail
	err = coll.FindOne(context.TODO(), filter).Decode(&dbres_3)
	ErrOK(err)
	dbres_3.App_List_num = 3
	//병합
	var will_send []Dj_jobs_detail
	will_send = append(will_send, dbres_1, dbres_2, dbres_3)
	will_send_json, _ := json.Marshal(will_send)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(will_send_json)
}

func Test2(w http.ResponseWriter, r *http.Request) { //applist를 보낼 때 모델이 될 예정
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

	coll := db.Database("dj_jobs").Collection("job_list")
	//쿼리1
	filter := bson.D{
		{"$and", bson.A{
			bson.D{{"사업장 주소", bson.D{{"$regex", "대구"}}}},
			bson.D{{"사업장 주소", bson.D{{"$regex", "북구"}}}},
			bson.D{{"필수부위", bson.D{{"$regex", "팔"}}}},
		},
		},
	}
	cursor, err := coll.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	var will_send []Dj_jobs_detail
	App_List_num := 0
	for cursor.Next(context.TODO()) {
		var dbres Dj_jobs_detail
		cursor.Decode(&dbres)
		dbres.App_List_num = App_List_num
		will_send = append(will_send, dbres)
		App_List_num++
	}
	ErrOK(err)

	will_send_json, _ := json.MarshalIndent(will_send, " ", "	")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(will_send_json)
}

/* 앤드 연산 예제
filter := bson.D{
   {"$and",
      bson.A{
         bson.D{{"rating", bson.D{{"$gt", 7}}}},
         bson.D{{"rating", bson.D{{"$lte", 10}}}},
      },
   },
}

*/
