package modules

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type job_detail_scrap struct {
	Scrap_List_num int                `json:"scrap_list_num"`
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
}

func PrintScrap(w http.ResponseWriter, r *http.Request) {
	if IsHeLogin(w, r) {
		var will_send []job_detail_scrap
		oid := SessionTO_oid(w, r)
		user_struct := OidTOuser_struct(oid)
		for i, v := range user_struct.ScrapList {
			dj_temp, err := OidTOjobDetail(v)
			ErrOK(err)

			//양식에 맞게 변환
			temp := job_detail_scrap{
				Scrap_List_num: i,
				ID:             dj_temp.ID,
				Duration:       dj_temp.Duration,
				CompanyName:    dj_temp.CompanyName,
				RecuritType:    dj_temp.RecuritType,
				RecuritShape:   dj_temp.RecuritShape,
				WageType:       dj_temp.WageType,
				Wage:           dj_temp.Wage,
				ComeType:       dj_temp.ComeType,
				RequireHistory: dj_temp.RequireHistory,
				RequireStudy:   dj_temp.RequireStudy,
				RelateMajor:    dj_temp.RelateMajor,
				RequireLicense: dj_temp.RequireLicense,
				Address:        dj_temp.Address,
				CompanyType:    dj_temp.CompanyType,
				ResponsiveIns:  dj_temp.ResponsiveIns,
				CreateAt:       dj_temp.CreateAt,
				Contact:        dj_temp.Contact,
				BodySpec:       dj_temp.BodySpec,
			}

			will_send = append(will_send, temp)
		}
		will_send_json, err := json.MarshalIndent(will_send, " ", "	")
		ErrOK(err)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(will_send_json)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		msg := map[string]string{"error": "Not LOGIN"}
		msg_json, _ := json.Marshal(msg)
		w.Write(msg_json)
	}

}
