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

func AuthHandler(w http.ResponseWriter, r *http.Request, password bool) {
	err := godotenv.Load()
	URI := os.Getenv("MONGODB_URI")
	if URI == "" {
		Critical(err)
	}
	form_email := r.FormValue("email")
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	Critical(err)
	defer func() {
		err := db.Disconnect(context.TODO())
		Critical(err)
	}()
	coll := db.Database("dj_users").Collection("users")
	filter := bson.D{{"email", form_email}}
	var dbres Dj_users_users
	err = coll.FindOne(context.TODO(), filter).Decode(&dbres)
	same_mail_not_found := func(err error) bool { //같은 email을 찾았는지 판별하는 anonymous 함수
		return err != nil
	}(err)
	if same_mail_not_found {
		fmt.Println("pseudo: 회원가입으로 이동하기", form_email)
		a := SmtpSender(form_email, true)
		fmt.Print(a)
	} else {
		fmt.Println("pseudo:", dbres.Email, "을 E-Mail로 로그인하기")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		redirect_msg := "<meta http-equiv=\"refresh\" content=\"0;url=/login/id/" + form_email + "\"></meta>"
		w.Write([]byte(redirect_msg))
	}
}
