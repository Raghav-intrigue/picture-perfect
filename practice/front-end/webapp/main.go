package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"clumio.com/roger/picture-perfect/webapp/controller"
	"clumio.com/roger/picture-perfect/webapp/middleware"
)

func main() {
	fmt.Println("Initialising...")

	templates := populateTemplates()
	controller.Startup(templates)

	fmt.Println("Server ready")

	http.ListenAndServe(":8000",
		&middleware.TimeoutMiddleware{
			Next: new(middleware.GzipMiddleware)})
}

func populateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _, fi := range files {

		tmpl := template.Must(layout.Clone())
		_, err = tmpl.ParseFiles(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}

		result[fi.Name()] = tmpl
	}
	return result
}
