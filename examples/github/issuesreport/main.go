package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"../github"
)

const templ = `Liczba znalezionych tematów {{.TotalCount}}:
{{range .Items}}----------------------------------------
Numer:        {{.Number}}
Użytkownik:   {{.User.Login}}
Tytuł:        {{.Title | printf "%.64s"}}
Utworzony:    {{.CreatedAt | daysAgo}} dni temu
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func noMust() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
