package myweb

import "html/template"

var Tmpl = template.Must(template.ParseGlob("../templates/*"))

type Greet struct {
	Greeting string
}
