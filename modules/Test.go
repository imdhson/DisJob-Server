package modules

import (
	"context"
	"encoding/json"
	"fmt"
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
	// 쿼리2
	filter = bson.D{{"사업장명", "용인시청"}}
	var dbres_2 Dj_jobs_detail
	err = coll.FindOne(context.TODO(), filter).Decode(&dbres_2)
	ErrOK(err)

	//쿼리3
	filter = bson.D{{"사업장명", "강원랜드(주)"}}
	var dbres_3 Dj_jobs_detail
	err = coll.FindOne(context.TODO(), filter).Decode(&dbres_3)
	ErrOK(err)

	//병합
	var will_send []Dj_jobs_detail
	will_send = append(will_send, dbres_1, dbres_2, dbres_3)
	will_send_json, _ := json.Marshal(will_send)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(will_send_json)
	fmt.Println(will_send[0].CompanyName)
	fmt.Println(will_send[1].CompanyName)
}
