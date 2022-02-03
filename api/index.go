package handler

import (
	"html/template"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	const tpl = `<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
<circle cx="50" cy="50" r="50"/>
</svg>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
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
