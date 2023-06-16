package main

import (
	our "disjob/modules"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func urlHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1:] //sv_urlpath에 유저가 어떤 url을 요청했는지 저장됨
	urlPath := strings.Split(url, "/")
	urlPath = append(urlPath, "", "", "") //인덱싱 out of range를 막기위해 빈 슬라이스  생성
	fmt.Println(urlPath)
	switch urlPath[0] {
	case "assets":
		our.AssetsHanlder(w, r, &url)
	default:
		our.HtmlHanlder(w, r, &urlPath)
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
