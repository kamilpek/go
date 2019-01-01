package main

import (
	"log"
	"os"

	"../github"
)

import "html/template"

var issueList = template.Must(template.New("issuelist").Parse(`
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
</head>
<h1>Liczba znalezionych tematów {{.TotalCount}}</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>Stan</th>
  <th>Użytkownik</th>
  <th>Tytuł</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
