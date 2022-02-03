package handler

import (
	"html/template"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	const tpl = `
<?xml version=\"1.0\" encoding=\"utf-8\" ?>
<svg width="640" height="480" xmlns="http://www.w3.org/2000/svg" xmlns:svg="http://www.w3.org/2000/svg">
	<!-- Created with SVG-edit - http://svg-edit.googlecode.com/ -->
	<g>
	<title>Layer 1</title>
	<rect id="svg_1" height="126" width="165" y="109" x="155" stroke-width="5" stroke="#000000" fill="#ff00ff"/>
	</g>
</svg>
 `

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("svg").Parse(tpl)
	check(err)

	type Data struct {
		Title       string
		Description string
		Percentage  int8
	}

	data := Data{
		Title:      "6% of 2022 has passed",
		Percentage: 10,
	}

	err = t.Execute(w, data)
	check(err)
}
