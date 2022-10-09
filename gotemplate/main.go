package main

import (
	"os"
	"text/template"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "name1f"}
	f2 := Friend{Fname: "name2f"}
	t := template.New("test")
	t = template.Must(t.Parse(
		`hello {{.UserName}}
{{ range .Emails }}
an email {{ . }}
{{- end }}
{{ with .Friends }}
{{- range . }}
my friend name is {{.Fname}}
{{- end }}
{{ end }}`))
	p := Person{UserName: "uname", Emails: []string{"email1", "email2"}, Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
