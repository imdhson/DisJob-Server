package main

import (
	"disjob/modules"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func urlHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1:] //sv_urlpath에 유저가 어떤 url을 요청했는지 저장됨
	urlPath := strings.Split(url, "/")
	urlPath = append(urlPath, "", "", "") //인덱싱 out of range를 막기위해 빈 슬라이스  생성
	fmt.Println(urlPath)
	switch urlPath[0] {
	case "login":
		if urlPath[1] == "auth" && urlPath[2] == "id" { //login/auth/id인 경우
			modules.AuthIDHandler(w, r) //id를 넘기는 모드
		} else if urlPath[1] == "auth" && urlPath[2] == "password" { //login/auth/password
			modules.AuthPWHandler(w, r) //비밀번호를 보내주는 상태
		} else if urlPath[1] == "auth" && urlPath[2] == "register" {
			modules.RegisterHandler(w, r)
		} else if urlPath[1] == "id" { //login/id 인경우
			modules.PWrequestHandler(w, r, &urlPath)
		} else {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			wwwfile, err := os.ReadFile("./www/login.html")
			modules.Critical(err)
			w.Write(wwwfile)
		}
	case "r":
		modules.RegisterPWrequestHandler(w, r, &urlPath)
	case "assets":
		modules.AssetsHanlder(w, r, &url)
	case "sample":
		modules.SampleAIList(w, r)
	case "sessiontest":
		oid := modules.SessionTO_oid(w, r)
		fmt.Println(oid)
		a := modules.OidTOuser_struct(oid)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(a.Email))

	default:
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		wwwfile, err := os.ReadFile("./www/error.html")
		modules.Critical(err)
		w.Write(wwwfile)
	}
}

func main() {
	const PORT int = 8080
	server := http.NewServeMux()
	server.Handle("/", http.HandlerFunc(urlHandler))
	fmt.Println("http://localhost:"+strconv.Itoa(PORT), "에서 요청을 기다리는 중:")
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), server)
	if err != nil { // http 서버 시작 중 문제 발생시
		log.Fatal(err)
		panic(err)
	}
}
