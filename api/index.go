package handler

import (
	"html/template"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>{{.Title}}</title>
			<style>
				.main {
						display: flex;
				}
				.container {
					width: 100%;
				}
				.bar {
					border: 5px solid #333333;
					background-color: #333333;
					width: 75%;
				}
				.progress {
					background-color: #00fe30;
					width: {{.Percentage}}%;
					height: 40px;
				}
				.title {
					color: #fff;
				}
		</style>
		</head>
		<body>
			<div class="main">
				<div class="container">
					<h1 class="">{{.Title}}</h1>
					<div class="bar">
						<div class="progress"></div>
					</div>
				</div>
			</div>
		</body>
	</html>
`

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

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = t.Execute(w, noItems)
	check(err)
}
