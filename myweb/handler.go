package myweb

import "net/http"

func ServeIndex(w http.ResponseWriter, r *http.Request) {

	Tmpl.ExecuteTemplate(w, "index", &Greet{Greeting: "yes Lord"})
}
