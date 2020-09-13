package Database

import "html/template"

var Main *template.Template
var Module *template.Template

func init() {
	Main = template.Must(template.ParseGlob("public/*.html"))
	Module = template.Must(template.ParseGlob("public/modules/*.html"))
}
