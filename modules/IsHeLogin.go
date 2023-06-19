package modules

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsHeLogin(w http.ResponseWriter, r *http.Request) bool {
	oid := SessionTO_oid(w, r)
	if oid != primitive.NilObjectID {
		return true
	} else {
		return false
	}
}
