package modules

import (
	"net/http"
	"os"
)

func HtmlHanlder(w http.ResponseWriter, r *http.Request, urlPath *[]string) {
	switch (*urlPath)[0] {
	case "login":
		if (*urlPath)[1] == "auth" && (*urlPath)[2] == "id" { //login/auth/id인 경우
			AuthHandler(w, r, false) //id를 넘기는 모드
		} else if (*urlPath)[1] == "auth" && (*urlPath)[2] == "password" { //login/auth/password
			AuthHandler(w, r, true) //비밀번호를 보내주는 상태
		} else if (*urlPath)[1] == "auth" && (*urlPath)[2] == "register" {
			RegisterHandler(w, r)
		} else if (*urlPath)[1] == "id" { //login/id 인경우
			PWrequestHandler(w, r, urlPath)
		} else {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			wwwfile, err := os.ReadFile("./www/login.html")
			Critical(err)
			w.Write(wwwfile)
		}
	case "r":
		RegisterPWrequestHandler(w, r, urlPath)
	default:
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		wwwfile, err := os.ReadFile("./www/error.html")
		Critical(err)
		w.Write(wwwfile)
	}
}
