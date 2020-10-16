package main

import (
	"fmt"
	// "io"
	"os"
	// "io/ioutil"
	models "github.com/therumbler/go-test/models"
	"text/template"
)

func checkErr(e error) {
	if e != nil {
		fmt.Println("ERROR: ", e)
		panic(e)
	}
}

func generate() {
	// var out io.Writer
	paths := []string{
		"templates/index.tmpl",
	}
	index := models.Index{Title: "This Is The Title"}
	indexTmpl, err := template.New("index.tmpl").ParseFiles(paths...)
	checkErr(err)
	err = indexTmpl.Execute(os.Stdout, index)
	checkErr(err)

	// fmt.Printf("Out file %q", out)
}

func main() {
	generate()
}
