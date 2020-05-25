package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"lss/webapp/src/lss/webapp/viewmodel"
)

// Struct that implements the handler interface (http.Handler)
// type myHandler struct {
// 	greeting string
// }

// func (mh myHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
// 	response.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
// }

// Templates
var templates = populateTemplates()

func main() {

	// Redirect /html/ links
	http.HandleFunc("/html/", func(response http.ResponseWriter, request *http.Request) {
		newURL := "http://" + request.Host + request.URL.Path[len("/html"):]
		if strings.HasSuffix(newURL, ".html") {
			newURL = newURL[:len(newURL)-5]
		}
		http.Redirect(response, request, newURL, http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		requestedFile := request.URL.Path[1:] // First character is '/'
		t := templates[requestedFile+".html"]
		if t != nil {

			// data context
			var context interface{}
			switch requestedFile {
			case "shop":
				context = viewmodel.NewShop()
			case "home":
				context = viewmodel.NewHome()
			default:
				context = viewmodel.NewBase()
			}

			err := t.Execute(response, context)
			if err != nil {
				log.Println(err)
			}
		} else {
			response.WriteHeader(http.StatusNotFound)
		}
	})

	/*
	** Built-in Handlers
	** NotFoundHandler:
	** RedirectHandler: redirects to another url
	** StripPrefix: passes to handler after stripping a prefix
	** TimoutHandler: passes to handler with a timelimit
	** FileServer: takes in a FileSystem interface ( ) // Dir type that implements FileSystem for the os
	 */

	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":8000", nil)
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

//region Defunct

func functionAsHandler() {

	// direct function that handles (no need for struct)
	http.HandleFunc("/public", func(response http.ResponseWriter, request *http.Request) {
		file, err := os.Open("public" + request.URL.Path)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer file.Close()

		// required to differentiate between plaintext and special types
		var contentType string
		switch {
		case strings.HasSuffix(request.URL.Path, "css"):
			contentType = "text/css"
		case strings.HasSuffix(request.URL.Path, "html"):
			contentType = "text/html"
		case strings.HasSuffix(request.URL.Path, "png"):
			contentType = "image/png"
		default:
			contentType = "text/plain"
		}
		response.Header().Add("Content-Type", contentType)

		io.Copy(response, file) // copies the file to the ResponseWriter
	})

	// Serve file method
	http.HandleFunc("/css", func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "public")
	})
}

func staticStringTemplate() {

	// use tick to escape characters
	templateString := `This is a dynamic title`
	// html/template safer as it escapes characters
	t, err := template.New("title").Parse(templateString)
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func readAllTemplates() *template.Template {
	const basePath = "templates"
	result := template.New("templates")
	template.Must(result.ParseGlob(basePath + "/*.html"))

	return result
}

//endregion
