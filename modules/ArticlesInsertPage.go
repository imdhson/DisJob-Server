package modules

import (
	"net/http"
	"os"
)

func ArticlesInsertPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	wwwfile, err := os.ReadFile("./www/articles_insert.html")
	Critical(err)

	form_action_url := "/articles/insert/submit"

	var i Vars_on_html
	i.Init()
	i.AddVar("form_action_url", form_action_url)
	wwwfile = i.VarsOnHTML(wwwfile)
	w.Write(wwwfile)
}
