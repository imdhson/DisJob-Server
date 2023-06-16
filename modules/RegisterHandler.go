package modules

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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
		msg := "비밀번호가 불일치 합니다. 이메일을 다시 눌러주세요."
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

	users_struct := Dj_users_users{Email: dbres.Email, Password: form_pw1, LastLogin: "now", Settings: Dj_users_users_settings{}}
	coll_dj_users := db.Database("dj_users").Collection("users")
	result, err := coll_dj_users.InsertOne(context.TODO(), users_struct)
	fmt.Println(result)
}
