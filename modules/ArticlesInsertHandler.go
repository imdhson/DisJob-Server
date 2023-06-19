package modules

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ArticlesInsertHandler(w http.ResponseWriter, r *http.Request) {
	if !IsHeLogin(w, r) {
		ErrHandler(w, r)
		return
	}
	user_oid := SessionTO_oid(w, r)

	title := r.FormValue("title")
	content := r.FormValue("content")

	var anon bool
	if r.FormValue("anonymous") == "on" {
		anon = true
	}
	err := godotenv.Load()
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
	articles_struct := Dj_board_articles{
		Djuserid: user_oid,
		CreateAt: time.Now(),
		Title:    title,
		Content:  content,
	}
	if anon {
		articles_struct.Djuserid = primitive.NilObjectID
	}
	_, err = coll.InsertOne(context.TODO(), articles_struct)
	if err != nil {
		ErrOK(err)
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		redirect_url := "/articles/"
		msg := "<meta http-equiv=\"refresh\" content=\"0;url=" + redirect_url + "\"></meta>"
		w.Write([]byte(msg))
	}

}
