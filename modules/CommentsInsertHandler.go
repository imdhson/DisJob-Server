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

func CommentsInsert(w http.ResponseWriter, r *http.Request, urlPath *[]string) { //comments/insert/@oid
	if !IsHeLogin(w, r) {
		ErrHandler(w, r)
		return
	}
	ARJB_oid_str := (*urlPath)[2] //article과 jobdetail에 모두 활용하는 oid
	ARJB_oid, err := primitive.ObjectIDFromHex(ARJB_oid_str)
	ErrOK(err)
	user_oid := SessionTO_oid(w, r)

	content := r.FormValue("content")

	var anon bool
	if r.FormValue("anonymous") == "on" {
		anon = true
	}
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
	coll := db.Database("dj_board").Collection("comments")
	comments_struct := Dj_board_comments{
		Djjobsid: ARJB_oid,
		Djuserid: user_oid,
		CreateAt: time.Now(),
		Content:  content,
	}
	if anon {
		comments_struct.Djuserid = primitive.NilObjectID
	}
	_, err = coll.InsertOne(context.TODO(), comments_struct)
	if err != nil {
		ErrOK(err)
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		redirect_url := "/comments/" + ARJB_oid_str
		msg := "<meta http-equiv=\"refresh\" content=\"0;url=" + redirect_url + "\"></meta>"
		w.Write([]byte(msg))
	}

}
