package modules

import (
	"encoding/json"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type job_scrap struct {
	Scrap_List_num int                `json:"scrap_list_num"`
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	CompanyName    string             `bson:"사업장명" json:"사업장명"`
	WageType       string             `bson:"임금형태" json:"임금형태"`
	Wage           int                `bson:"임금" json:"임금"`
	Address        string             `bson:"사업장 주소" json:"사업장 주소"`
	RecuritShape   string             `bson:"고용형태" json:"고용형태"`
}

func PrintScrap(w http.ResponseWriter, r *http.Request) {
	if IsHeLogin(w, r) {
		var will_send []job_scrap
		oid := SessionTO_oid(w, r)
		user_struct := OidTOuser_struct(oid)
		for i, v := range user_struct.ScrapList {
			dj_temp, err := OidTOjobDetail(v)
			ErrOK(err)
			tmp_address := strings.Split(dj_temp.Address, " ") //위치 앞에만 잘라서 보냄
			tmp_address1 := tmp_address[0] + " " + tmp_address[1]

			temp := job_scrap{
				Scrap_List_num: i,
				ID:             dj_temp.ID,
				CompanyName:    dj_temp.CompanyName,
				WageType:       dj_temp.WageType,
				Wage:           dj_temp.Wage,
				Address:        tmp_address1,
				RecuritShape:   dj_temp.RecuritShape,
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
