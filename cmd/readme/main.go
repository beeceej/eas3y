package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func buildReadme() {
	var (
		fexample *os.File
		freadme  *os.File
		b        []byte
		err      error
	)

	if fexample, err = os.Open(filepath.Join("..", "example", "main.go")); err != nil {
		panic(err.Error())
	}
	if freadme, err = os.OpenFile(filepath.Join("..", "..", "README.md"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777); err != nil {
		panic(err.Error())
	}

	if b, err = ioutil.ReadAll(fexample); err != nil {
		panic(err.Error())
	}

	tmpl := template.Must(template.New("readme.tmpl").ParseFiles("readme.tmpl"))
	if err != nil {
		panic(err.Error())
	}
	if err = tmpl.Execute(freadme, string(b)); err != nil {
		panic(err.Error())
	}
}

func main() {
	buildReadme()
}
