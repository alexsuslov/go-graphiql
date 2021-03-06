// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate

package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
	"time"
)
const url = "https://raw.githubusercontent.com/graphql/graphiql/master/example/index.html"



func main(){

	bytes, err := ioutil.ReadFile("html/index.html")
	//bytes, err := readFile(url)
	if err != nil {
		panic(err)
	}

	//log.Println(string(bytes))

	f, err := os.Create("html/html.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = packageTemplate.Execute(f, struct {
		Timestamp time.Time
		URL string
		CONTENT   string
		BQ string
	}{
		Timestamp: time.Now(),
		URL : url,
		CONTENT:   string(bytes),
		BQ : "`",
	})

	if err != nil {
		panic(err)
	}
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
// {{ .URL }}

package html

var Content = {{ .BQ }}{{ .CONTENT }}{{ .BQ }}
`))

func readFile(Url string)([]byte, error){
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}