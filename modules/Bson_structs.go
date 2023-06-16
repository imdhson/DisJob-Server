package modules

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dj_users_users struct {
	ID        primitive.ObjectID      `bson:"_id,omitempty"`
	Email     string                  `bson:"email,omitempty"`
	Password  string                  `bson:"password,omitempty"`
	LastLogin string                  `bson:"lastLogin,omitempty"`
	ScrapList []interface{}           `bson:"scrapList,omitempty"`
	Settings  Dj_users_users_settings `bson:"settings,omitempty"`
}

type Dj_users_scraplist struct {
}

type Dj_users_users_settings struct {
	Loc  string `bson:"loc"`
	Type string `bson:"type"`
}

type Dj_users_registration struct {
	ID           primitive.ObjectID `bson:"_id"`
	Email        string             `bson:"email"`
	VerifyNumber string             `bson:"verifyNumber"`
}
