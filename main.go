package main

import (
	"disjob/modules"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 로그 기록
	// 로그 파일 생성
	log_f, log_err := os.OpenFile("last.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	modules.Critical(log_err)
	defer log_f.Close()

	// 로그 출력 설정
	log.SetOutput(io.MultiWriter(os.Stdout, log_f))

	const PORT int = 443
	server := http.NewServeMux()
	server.Handle("/", http.HandlerFunc(urlHandler))
	log.Println("http://localhost:"+strconv.Itoa(PORT), "에서 요청을 기다리는 중:")

	go http.ListenAndServe(":80", server) //일반 http 전송 디버그용

	err := http.ListenAndServeTLS(":"+strconv.Itoa(PORT), "openssl/certificate.crt", "openssl/private.key", server)
	modules.Critical(err)
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1:] //sv_urlpath에 유저가 어떤 url을 요청했는지 저장됨
	urlPath := strings.Split(url, "/")
	urlPath = append(urlPath, "", "", "") //인덱싱 out of range를 막기위해 빈 슬라이스  생성
	log.Printf("%v/%v", modules.GetIP(r), url)
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
		} else if urlPath[1] == "logout" { //login/logout 인경우
			modules.LogoutHandler(w, r)
		} else {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			wwwfile, err := os.ReadFile("./www/login.html")
			modules.Critical(err)
			w.Write(wwwfile)
		}
	case "r":
		modules.RegisterPWrequestHandler(w, r, &urlPath)
	case "comments":
		if urlPath[1] == "insert" {
			modules.CommentsInsert(w, r, &urlPath)
		} else {
			//modules.CommentsView(w, r, &urlPath)
			modules.ErrHandler(w, r)
		}
	case "articles":
		if urlPath[1] == "insert" && urlPath[2] == "" { //삽입 모드
			modules.ArticlesInsertPage(w, r)
		} else if urlPath[1] == "insert" && urlPath[2] == "submit" { //삽입 제출 모드
			modules.ArticlesInsertHandler(w, r)
		} else if urlPath[1] == "" { //게시글 제목 뷰 모드
			modules.ArticlesView(w, r)
		} else {
			modules.ArticlesDetailHandler(w, r, &urlPath)
		}

	case "assets":
		modules.AssetsHanlder(w, r, &url)

	case "users":
		if urlPath[1] == "settings" && urlPath[2] == "submit" {
			modules.SettingsChangeHandler(w, r)
		} else {
			modules.ErrHandler(w, r)
		}
	case "ailist":
		modules.AIListSender(w, r)
	case "sessiontest":
		modules.PrintSession(w, r)
	case "session":
		modules.PrintSession(w, r)
	case "test":
		modules.SampleAIList(w, r)
	case "test2":
		modules.Test2(w, r)
	case "test3":
		modules.Test3(w, r)
	default:
		modules.ErrHandler(w, r)
	}
}
