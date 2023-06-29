package modules

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PrintJobDetail(w http.ResponseWriter, r *http.Request, urlPath *[]string) {
	oid_hex := (*urlPath)[1]
	oid, err := primitive.ObjectIDFromHex(oid_hex)
	ErrOK(err)
	temp, err := OidTOjobDetail(oid)
	if err != nil { //job을 찾지 못하였을 때
		temp := map[string]string{"사업장명": "찾지 못함"}
		temp2, err := json.MarshalIndent(temp, " ", "	")
		ErrOK(err)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(temp2)
	} else {
		temp2, err := json.MarshalIndent(temp, " ", "	")
		ErrOK(err)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(temp2)
	}
}
