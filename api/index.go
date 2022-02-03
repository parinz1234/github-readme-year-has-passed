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

	const tpl = `{{.Top}}<svg viewBox="0 0 400 130" width="400" height="130" xmlns="http://www.w3.org/2000/svg" xmlns:svg="http://www.w3.org/2000/svg"><rect width="380" height="80" style="stroke: rgb(0, 0, 0); fill: rgb(51, 51, 51); fill-rule: nonzero;" data-bx-origin="-0.038 -0.658333" x="0" y="50"></rect>
<text style="white-space: pre; fill: rgb(51, 51, 51); font-family: Arial, sans-serif; font-size: 17.2px;" x="258.918" y="49.952"> </text><g></g><text style="fill: rgb(51, 51, 51); font-family: Arial, sans-serif; font-size: 19.2px; white-space: pre;" data-bx-origin="-0.076015 -0.947915" transform="matrix(1.245747, 0, 0, 0.950186, -4.914937, 2.73884)"><tspan x="0" y="35.375"">6% of 2022 has passed</tspan></text><rect x="30" y="60" width="360" height="60" style="stroke: rgb(0, 0, 0); fill: rgb(0, 253, 49);"></rect></svg>
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
	w.Header().Set("Content-Type", "image/svg+xml")
	err = t.ExecuteTemplate(w, "svg", &d)
	check(err)
}
