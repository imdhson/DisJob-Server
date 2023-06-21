package modules

import (
	"encoding/json"
	"net/http"
)

func PrintSession(w http.ResponseWriter, r *http.Request) {
	if IsHeLogin(w, r) {
		oid := SessionTO_oid(w, r)
		temp := OidTOuser_struct(oid)
		temp2, _ := json.MarshalIndent(temp, " ", "	")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(temp2)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		msg := map[string]string{"email": "Not LOGIN"}
		msg_json, _ := json.Marshal(msg)
		w.Write(msg_json)
	}

}
