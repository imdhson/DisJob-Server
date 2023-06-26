package modules

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ArticlesView(w http.ResponseWriter, r *http.Request) {
	if !IsHeLogin(w, r) {
		ErrHandler(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	wwwfile, err := os.ReadFile("./www/articles.html")
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
	opts := options.Find().SetSort(bson.D{{"createAt", 1}})
	cursor, err := coll.Find(context.TODO(), bson.D{}, opts)
	ErrOK(err)

	var will_send []Dj_board_articles
	for cursor.Next(context.TODO()) {
		var dbres Dj_board_articles
		err := cursor.Decode(&dbres)
		ErrOK(err)
		will_send = append(will_send, dbres)
	}
	ErrOK(err)

	var title_msg string

	for _, v := range will_send {
		compare_time := time.Since(v.CreateAt).String()
		compare_time_splited := strings.Split(compare_time, "m") //이후로 무시하기 위함
		useremail := OidTOuser_struct(v.Djuserid).Email
		article_url := "/articles/" + v.ID.Hex()
		title_msg +=
			"<a href=\"" + article_url + "\"" + // /articles/oid 형식으로 전달되게됨
				"<li><span class=\"comment-content\">" +
				v.Title +
				"</span></a><span class=\"comment-writer\">" +
				useremail + "(이)가 " + compare_time_splited[0] + "분 전 작성</span></li>"
	}
	htmlmodify.AddVar("title_msg", title_msg)
	html_modified := htmlmodify.VarsOnHTML(wwwfile)
	w.Write(html_modified)
}
