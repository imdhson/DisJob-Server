package modules

import (
	"net/http"
	"os"
)

func PWrequestHandler(w http.ResponseWriter, r *http.Request, urlPath *[]string) {
	url_email := (*urlPath)[2]

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	wwwfile, err := os.ReadFile("./www/login_pw.html")
	Critical(err)
	//wwwfile, err := template.ParseFiles("./www/login_pw.html")

	var Vars Vars_on_html
	Vars.Init()
	Vars.AddVar("url_email", url_email)
	Vars.Display()

	w.Write(Vars.VarsOnHTML(wwwfile))
	//	fmt.Println(string(wwwfile))

}
