package modules

import (
	"context"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterPWrequestHandler(w http.ResponseWriter, r *http.Request, urlPath *[]string) {
	url_email := (*urlPath)[1]
	url_key := (*urlPath)[2]
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
	coll := db.Database("dj_users").Collection("registration")
	filter := bson.D{{"email", url_email}}
	var dbres Dj_users_registration
	err = coll.FindOne(context.TODO(), filter).Decode(&dbres)

	ErrOK(err)
	if dbres.VerifyNumber == url_key {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		wwwfile, err := os.ReadFile("./www/register.html")
		Critical(err)
		var i Vars_on_html
		i.Init()
		i.AddVar("url_email", url_email)
		i.AddVar("url_key", url_key)
		i.VarsOnHTML(wwwfile)
		w.Write(i.VarsOnHTML(wwwfile))
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		redirect_msg := "<meta http-equiv=\"refresh\" content=\"0;url=/login\">"
		w.Write([]byte(redirect_msg))
	}

}
