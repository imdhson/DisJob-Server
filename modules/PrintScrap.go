package modules

import (
	"encoding/json"
	"net/http"
)

func PrintScrap(w http.ResponseWriter, r *http.Request) {
	if IsHeLogin(w, r) {
		var will_send Dj_jobs_detail_s
		oid := SessionTO_oid(w, r)
		user_struct := OidTOuser_struct(oid)
		for _, v := range user_struct.ScrapList {
			temp, err := OidTOjobDetail(v)
			ErrOK(err)
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
		msg := map[string]string{"email": "Not LOGIN"}
		msg_json, _ := json.Marshal(msg)
		w.Write(msg_json)
	}

}
