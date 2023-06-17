package modules

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	URI := os.Getenv("MONGODB_URI")
	if URI == "" {
		Critical(err)
	}
	form_key := r.FormValue("verifyNumber")
	form_pw1 := r.FormValue("password1")
	form_pw2 := r.FormValue("password2")
	if form_pw1 != form_pw2 {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		msg := "<script>alert(\"비밀번호가 불일치합니다. 다시 입력해주세요.\")</script><meta http-equiv=\"refresh\" content=\"0;url=/login/\"></meta>"
		w.Write([]byte(msg))
	}

	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	Critical(err)
	defer func() {
		err := db.Disconnect(context.TODO())
		Critical(err)
	}()
	coll_dj_registration := db.Database("dj_users").Collection("registration")
	var dbres Dj_users_users
	filter_for_key_email_search := bson.D{{"verifyNumber", string(form_key)}}
	err = coll_dj_registration.FindOne(context.TODO(), filter_for_key_email_search).Decode(&dbres)

	now_time := time.Now()
	users_struct := Dj_users_users{
		Email: dbres.Email, Password: form_pw1,
		LastLogin: now_time, ScrapList: []primitive.ObjectID{{}}, // primitive.NewObjectID 로 나중에 Push 가능
		Settings: Dj_users_users_settings{},
	}
	coll_dj_users := db.Database("dj_users").Collection("users")
	result, err := coll_dj_users.InsertOne(context.TODO(), users_struct)
	fmt.Println(result)
}
