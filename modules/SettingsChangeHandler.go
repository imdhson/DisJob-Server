package modules

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Receive_settings struct {
	Loc  string `json:"loc"`
	Type string `json:"type"`
}

func SettingsChangeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // 모든 도메인에 대해 허용
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "POST" {
		//json을 수신하여 settings_struct에 저장
		var settings_struct Receive_settings
		body, err := io.ReadAll(r.Body)
		fmt.Println("json 수신 원본: ", string(body))
		ErrOK(err)
		err = json.Unmarshal(body, &settings_struct)
		ErrOK(err)
		defer r.Body.Close()
		if !IsHeLogin(w, r) { //로그인 안했으면
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			redirect_msg := "<script>alert(\"로그인 후 다시 시도.\")</script><meta http-equiv=\"refresh\" content=\"0;url=/login/\"></meta>"
			w.Write([]byte(redirect_msg))
			return
		}
		user_oid := SessionTO_oid(w, r)
		godotenv.Load()
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
		coll := db.Database("dj_users").Collection("users")

		filter := bson.D{{"_id", user_oid}}
		//update := bson.D{{"$set", bson.D{{"settings", bson.D{{"loc", settings_struct.Loc}}, bson.D{{"type", settings_struct.Type}}}}}}
		update := bson.D{
			{"$set", bson.D{{"settings.loc", settings_struct.Loc}, {"settings.type", settings_struct.Type}}}}

		_, err = coll.UpdateOne(context.TODO(), filter, update)
		ErrOK(err)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("good"))
	}

}
