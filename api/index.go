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

	const tpl = `{{.Top}}<svg viewBox="0 0 380 110" width="380" height="110" xmlns="http://www.w3.org/2000/svg" xmlns:svg="http://www.w3.org/2000/svg"><defs><style data-bx-fonts="Kanit">@import url(https://fonts.googleapis.com/css2?family=Kanit%3Aital%2Cwght%400%2C100%3B0%2C200%3B0%2C300%3B0%2C400%3B0%2C500%3B0%2C600%3B0%2C700%3B0%2C800%3B0%2C900%3B1%2C100%3B1%2C200%3B1%2C300%3B1%2C400%3B1%2C500%3B1%2C600%3B1%2C700%3B1%2C800%3B1%2C900&amp;display=swap);</style></defs><rect width="380" height="80" style="stroke: rgb(0, 0, 0); fill: rgb(51, 51, 51); fill-rule: nonzero;" data-bx-origin="-0.038 -0.658333" y="30"></rect><text style="white-space: pre; fill: rgb(51, 51, 51); font-family: Arial, sans-serif; font-size: 17.2px;" x="258.918" y="49.952"> </text><g></g><rect x="10" y="40" width="360" height="60" style="stroke: rgb(0, 0, 0); fill: rgb(0, 253, 49);"></rect><text style="fill: rgb(51, 51, 51); font-family: Kanit; font-size: 17.2px; font-weight: 600; white-space: pre;" x="0" y="19.065" transform="matrix(1, 0, 0, 0.967857, 0, -0.305795)">6% of 2022 has passed</text></svg>
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
