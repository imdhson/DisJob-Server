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

func ArticlesDetailHandler(w http.ResponseWriter, r *http.Request, urlPath *[]string) {
	if !IsHeLogin(w, r) {
		ErrHandler(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	wwwfile, err := os.ReadFile("./www/articles_detail.html")
	Critical(err)
	var htmlmodify Vars_on_html
	htmlmodify.Init()

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
	coll := db.Database("dj_board").Collection("articles")
	var dbres Dj_board_articles
	url_oid := (*urlPath)[1]
	oid, err := primitive.ObjectIDFromHex(url_oid)
	ErrOK(err)
	err = coll.FindOne(context.TODO(), bson.D{{"_id", oid}}).Decode(&dbres)
	articlemode := false
	if err == nil {
		articlemode = true
	}

	var title_msg string
	useremail := OidTOuser_struct(dbres.Djuserid).Email
	if dbres.Djuserid == primitive.NilObjectID {
		useremail = "익명의 유저"
	} else {
		useremail, _, _ = strings.Cut(useremail, "@")
	}
	compare_time := time.Since(dbres.CreateAt).String()
	compare_time, _, _ = strings.Cut(compare_time, "m") // m 이후로 무시하기 위함
	if strings.Contains(compare_time, ".") {            //1분 미만이면 방금이라고 표기
		compare_time = "방금 "
	} else {
		compare_time += "분 " //1분 이상이면 숫자+분
	}
	compare_time = strings.ReplaceAll(compare_time, "h", "시간")
	compare_time = strings.ReplaceAll(compare_time, "d", "일")
	title_msg += "<ul id=\"article-list\"><li class=\"article-detail\"><span class=\"comment-content-detail\">" + dbres.Title + "</span>" + "<span class=\"comment-writer\">" +
		useremail + "(이)가 " + compare_time + "전 작성</span>"
	title_msg += "<span class=\"article-content\">" + dbres.Content + "</span>"
	title_msg += "</li></ul>"

	coll_comments := db.Database("dj_board").Collection("comments")
	var comments_struct []Dj_board_comments

	filter := bson.D{{"dj_jobs_id", oid}}
	cursor, err := coll_comments.Find(context.TODO(), filter)
	ErrOK(err)
	for cursor.Next(context.TODO()) {
		var dbres_comments Dj_board_comments
		cursor.Decode(&dbres_comments)
		comments_struct = append(comments_struct, dbres_comments)
	}

	htmlmodify.AddVar("form_action_url", "/comments/insert/"+url_oid)
	var comments_msg string
	for _, v := range comments_struct {
		useremail := OidTOuser_struct(v.Djuserid).Email
		if v.Djuserid == primitive.NilObjectID {
			useremail = "익명의 유저"
		} else {
			useremail, _, _ = strings.Cut(useremail, "@")
		}
		compare_time := time.Since(v.CreateAt).String()
		compare_time, _, _ = strings.Cut(compare_time, "m") // m 이후로 무시하기 위함
		if strings.Contains(compare_time, ".") {            //1분 미만이면 방금이라고 표기
			compare_time = "방금 "
		} else {
			compare_time += "분 " //1분 이상이면 숫자+분
		}
		compare_time = strings.ReplaceAll(compare_time, "h", "시간")
		compare_time = strings.ReplaceAll(compare_time, "d", "일")
		comments_msg +=
			"<li><span class=\"comment-content\">" +
				v.Content +
				"</span><span class=\"comment-writer\">" +
				useremail + "(이)가 " + compare_time + "전 작성</span></li>"
	}
	//버튼 관련
	var button_msg string
	if !articlemode { //article이 아니기때문에 비움
		title_msg = ""
		button_msg = ""
	} else {
		button_msg = "<button style=\"position: fixed; right: 30px; bottom: 90px; font-size: 55px; background-color: rgb(202, 240, 255); border-radius: 50%; box-shadow: 0px 10px 20px rgb(170, 170, 170); width: 70px; height: 70px;border-color: transparent;\" onclick=\"location.href='/articles/insert'\">✍</button><button style=\"position: fixed; left: 30px; bottom: 90px; font-size: 40px; background-color: rgb(202, 240, 255); border-radius: 50%; box-shadow: 0px 10px 20px rgb(170, 170, 170); width: 70px; height: 70px;border-color: transparent;\" onclick=\"location.href='/articles'\"><</button>"

	}
	htmlmodify.AddVar("button_msg", button_msg)
	htmlmodify.AddVar("title_msg", title_msg)
	htmlmodify.AddVar("comments_msg", comments_msg)
	html_modified := htmlmodify.VarsOnHTML(wwwfile)
	w.Write(html_modified)
}
