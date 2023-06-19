package modules

import (
	"context"
	"net/http"
	"os"
	"strings"

	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CommentsView(w http.ResponseWriter, r *http.Request, urlPath *[]string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	wwwfile, err := os.ReadFile("./www/comments.html")
	Critical(err)
	var htmlmodify Vars_on_html
	htmlmodify.Init()

	ARJB_oid_str := (*urlPath)[1] //article과 jobdetail에 모두 활용하는 oid
	ARJB_oid, err := primitive.ObjectIDFromHex(ARJB_oid_str)
	ErrOK(err)
	err = godotenv.Load()
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
	coll := db.Database("dj_board").Collection("comments")
	var will_send []Dj_board_comments

	filter := bson.D{{"dj_jobs_id", ARJB_oid}}
	cursor, err := coll.Find(context.TODO(), filter)
	ErrOK(err)
	for cursor.Next(context.TODO()) {
		var dbres Dj_board_comments
		cursor.Decode(&dbres)
		will_send = append(will_send, dbres)
	}
	ErrOK(err)

	htmlmodify.AddVar("form_action_url", "/comments/insert/"+ARJB_oid_str)
	var comments_msg string
	for _, v := range will_send {
		useremail := OidTOuser_struct(v.Djuserid).Email
		if v.Djuserid == primitive.NilObjectID {
			useremail = "익명의 유저"
		}
		compare_time := time.Since(v.CreateAt).String()
		compare_time_splited := strings.Split(compare_time, ".") //이후로 무시하기 위함
		comments_msg +=
			"<li><span class=\"comment-content\">" +
				v.Content +
				"</span><span class=\"comment-writer\">" +
				useremail + "(이)가 " + compare_time_splited[0] + "초 전 작성</span></li>"
	}
	htmlmodify.AddVar("comments_msg", comments_msg)
	html_modified := htmlmodify.VarsOnHTML(wwwfile)
	w.Write(html_modified)
}
