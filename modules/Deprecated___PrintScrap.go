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

func will_send_append_scrap(dbres *job_scrap, input *[]job_scrap) bool {
	var tmp bool = false
	for _, v := range *input {
		if v.ID == dbres.ID { //will_send에 이미 포함 되어있는 데이터일때
			tmp = true
			return false
		} else { //포함되지 않았을 때 dbres를 append함
			tmp = false
		}
	}
	if !tmp { //포함되지 않았을 때 dbres를 append함
		//log.Println("어펜드 시도", dbres.ID)
		*input = append(*input, *dbres)
		return true
	}
	return false
}
func PrintScrap(w http.ResponseWriter, r *http.Request) {
	if IsHeLogin(w, r) {
		var will_send []job_scrap
		oid := SessionTO_oid(w, r)
		user_struct := OidTOuser_struct(oid)
		scrap_list_num := 0
		for _, v := range user_struct.ScrapList {

			dj_temp, err := OidTOjobDetail(v)
			if err != nil {
				ErrOK(err)
				continue
			}

			tmp_address := strings.Split(dj_temp.Address, " ") //위치 앞에만 잘라서 보냄
			tmp_address1 := tmp_address[0] + " " + tmp_address[1]

			temp := job_scrap{
				Scrap_List_num: scrap_list_num,
				ID:             dj_temp.ID,
				CompanyName:    dj_temp.CompanyName,
				WageType:       dj_temp.WageType,
				Wage:           dj_temp.Wage,
				Address:        tmp_address1,
				RecuritShape:   dj_temp.RecuritShape,
			}
			//시급으로 환산
			switch temp.WageType {
			case "일급":
				temp.Wage = temp.Wage / 8
				temp.WageType = "환산 시급"
			case "월급":
				temp.Wage = temp.Wage / (5 * 4 * 8)
				temp.WageType = "환산 시급"
			case "연봉":
				temp.Wage = temp.Wage / (12 * 5 * 4 * 8)
				temp.WageType = "환산 시급"
			}

			if will_send_append_scrap(&temp, &will_send) { // 중복하지않고 append 성공하면 리스트번호 1 추가
				scrap_list_num++
			}

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
