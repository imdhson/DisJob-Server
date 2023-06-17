package modules

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AuthPWHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	URI := os.Getenv("MONGODB_URI")
	if URI == "" {
		Critical(err)
	}
	form_email := r.FormValue("email")
	form_pw := r.FormValue("password")
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	Critical(err)
	defer func() {
		err := db.Disconnect(context.TODO())
		Critical(err)
	}()
	coll := db.Database("dj_users").Collection("users")
	filter := bson.D{{"email", form_email}, {"password", form_pw}}
	var dbres Dj_users_users
	err = coll.FindOne(context.TODO(), filter).Decode(&dbres)
	if err != nil {
		fmt.Println("ID, 비밀번호 매칭 실패")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		redirect_msg := "<script>alert(\"로그인 실패\")</script><meta http-equiv=\"refresh\" content=\"0;url=/login/id/" + form_email + "\"></meta>" //다시 원래 pwrequst
		w.Write([]byte(redirect_msg))
	} else {
		//로그인이 성공함
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmp := fmt.Sprintln("Login 성공 당신의 아이디는:", dbres.Email, "\n모든 내용:", dbres)
		sessionkey := rand.Int31()
		cookieid := http.Cookie{
			Name:     "dj_session",
			Value:    strconv.Itoa(int(sessionkey)),
			HttpOnly: true,
		}
		http.SetCookie(w, &cookieid)
		//db에 세션키 저장

		session_struct := Dj_user_session{
			Djuserid: dbres.ID,
			Session:  int(sessionkey),
			CreateAt: time.Now(),
		}
		coll_dj_users := db.Database("dj_users").Collection("sessions")
		result, err := coll_dj_users.InsertOne(context.TODO(), session_struct)
		ErrOK(err)
		fmt.Println(result.InsertedID)
		w.Write([]byte(tmp))
	}
}
