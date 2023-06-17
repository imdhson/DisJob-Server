package modules

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dj_users_users struct {
	ID        primitive.ObjectID      `bson:"_id,omitempty"`
	Email     string                  `bson:"email"`
	Password  string                  `bson:"password"`
	LastLogin time.Time               `bson:"lastLogin"`
	ScrapList []primitive.ObjectID    `bson:"scrapList"`
	Settings  Dj_users_users_settings `bson:"settings"`
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

type Dj_jobs_detail struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Duration       string             `bson:"구인신청일자"`
	CompanyName    string             `bson:"사업장명"`
	RecuritType    string             `bson:"모집직종"`
	RecuritShape   string             `bson:"고용형태"`
	WageType       string             `bson:"임금형태"`
	Wage           int                `bson:"임금"`
	ComeType       string             `bson:"입사형태"`
	RequireHistory string             `bson:"요구경력"`
	RequireStudy   string             `bson:"요구학력"`
	RelateMajor    string             `bson:"전공계열"`
	RequireLicense string             `bson:"요구자격증"`
	Address        string             `bson:"사업장 주소"`
	CompanyType    string             `bson:"기업형태"`
	ResponsiveIns  string             `bson:"담당기관"`
	CreateAt       string             `bson:"등록일"`
	Contact        string             `bson:"연락처"`
}
