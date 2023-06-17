package modules

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dj_users_users struct {
	ID        primitive.ObjectID      `bson:"_id,omitempty"`
	Email     string                  `bson:"email,omitempty"`
	Password  string                  `bson:"password,omitempty"`
	LastLogin time.Time               `bson:"lastLogin,omitempty"`
	ScrapList []primitive.ObjectID    `bson:"scrapList,omitempty"`
	Settings  Dj_users_users_settings `bson:"settings,omitempty"`
}

type Dj_users_users_settings struct {
	Loc  string `bson:"loc"`
	Type string `bson:"type"`
}

type Dj_users_registration struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Email        string             `bson:"email"`
	VerifyNumber string             `bson:"verifyNumber"`
	CreateAt     time.Time          `bson:"createAt"`
}

type Dj_user_session struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Djuserid primitive.ObjectID `bson:"dj_user_id"`
	Session  int                `bson:"dj_session"`
	CreateAt time.Time          `bson:"createAt"`
}
