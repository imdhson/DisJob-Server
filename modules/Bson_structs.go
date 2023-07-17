package modules

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dj_users_users struct {
	ID        primitive.ObjectID      `bson:"_id,omitempty"`
	Email     string                  `bson:"email"`
	Password  [64]byte                `bson:"password"`
	LastLogin time.Time               `bson:"lastLogin"`
	ScrapList []primitive.ObjectID    `bson:"scrapList"`
	Settings  Dj_users_users_settings `bson:"settings"`
}

type Dj_users_users_settings struct {
	Loc   string `bson:"loc"`
	Type1 string `bson:"type1"`
	Type2 string `bson:"type2"`
	Type3 string `bson:"type3"`
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
	AI_List_num    int
	AI_List_score  int
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Duration       string             `bson:"구인신청일자" json:"구인신청일자"`
	CompanyName    string             `bson:"사업장명" json:"사업장명"`
	RecuritType    string             `bson:"모집직종" json:"모집직종"`
	RecuritShape   string             `bson:"고용형태" json:"고용형태"`
	WageType       string             `bson:"임금형태" json:"임금형태"`
	Wage           int                `bson:"임금" json:"임금"`
	ComeType       string             `bson:"입사형태" json:"입사형태"`
	RequireHistory string             `bson:"요구경력" json:"요구경력"`
	RequireStudy   string             `bson:"요구학력" json:"요구학력"`
	RelateMajor    string             `bson:"전공계열" json:"전공계열"`
	RequireLicense string             `bson:"요구자격증" json:"요구자격증"`
	Address        string             `bson:"사업장 주소" json:"사업장 주소"`
	CompanyType    string             `bson:"기업형태" json:"기업형태"`
	ResponsiveIns  string             `bson:"담당기관" json:"담당기관"`
	CreateAt       string             `bson:"등록일" json:"등록일"`
	Contact        string             `bson:"연락처" json:"연락처"`
	BodySpec       string             `bson:"필수부위" json:"필수부위"`
	ScrapCount     int                `bson:"scrapCount" json:"scrapCount"`
}
type Dj_jobs_detail_s []Dj_jobs_detail
type Dj_jobs_refined struct {
	AI_List_num  int
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	CompanyName  string             `bson:"사업장명" json:"사업장명"`
	WageType     string             `bson:"임금형태" json:"임금형태"`
	Wage         int                `bson:"임금" json:"임금"`
	Address      string             `bson:"사업장 주소" json:"사업장 주소"`
	RecuritShape string             `bson:"고용형태" json:"고용형태"`
}

func (a Dj_jobs_detail_s) Len() int {
	return len(a)
}

func (a Dj_jobs_detail_s) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Dj_jobs_detail_s) Less(i int, j int) bool {
	return a[i].AI_List_score < a[j].AI_List_score
}

type Dj_board_comments struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Djjobsid primitive.ObjectID `bson:"dj_jobs_id"`
	Djuserid primitive.ObjectID `bson:"dj_user_id"`
	CreateAt time.Time          `bson:"createAt"`
	Content  string             `bson:"content"`
	GenbyAI  bool               `bson:"genbyAI"`
}
type Dj_board_articles struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Djuserid primitive.ObjectID `bson:"dj_user_id"`
	CreateAt time.Time          `bson:"createAt"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}
type Dj_jobs_typeavt struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Type         string             `bson:"종류"`
	Availability string             `bson:"가능 부위"`
}
