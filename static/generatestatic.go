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
	err := os.Mkdir("public", 0755)
	if err != nil {
		fmt.Println("public folder exists already")
	} else {
		fmt.Println("public folder created")
	}
	out, err := os.Create("public/index.html")
	checkErr(err)
	paths := []string{
		"templates/index.tmpl",
	}
	index := models.Index{Title: "This Is The Title 👋"}
	indexTmpl, err := template.New("index.tmpl").ParseFiles(paths...)
	checkErr(err)
	// err = indexTmpl.Execute(os.Stdout, index)
	err = indexTmpl.Execute(out, index)
	checkErr(err)
	out.Close()
	// fmt.Printf("Out file %q", out)
}

func main() {
	generate()
}
