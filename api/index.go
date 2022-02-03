package handler

import (
	"html/template"
	"log"
	"net/http"
)

type data struct {
	Top   template.HTML
	Color string
}

func Handler(w http.ResponseWriter, r *http.Request) {

	const tpl = `{{.Top}}<svg width="640" height="480" xmlns="http://www.w3.org/2000/svg" xmlns:svg="http://www.w3.org/2000/svg"><g><title>Layer 1</title><rect id="svg_1" height="126" width="165" y="109" x="155" stroke-width="5" stroke="#000000" fill="{{.Color}}"/></g></svg>
	`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	t, err := template.New("svg").Parse(tpl)
	check(err)

	var d data
	d.Top = "<?xml version=\"1.0\" encoding=\"utf-8\" ?>"
	d.Color = "#ff00ff"
	err = t.ExecuteTemplate(w, "svg", &d)
	check(err)
}
