package controller

import (
	"html/template"
	"net/http"

	"clumio.com/roger/picture-perfect/webapp/viewmodel"
)

type home struct {
	homeTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	w.Header().Add("Content-Type", "text/html")
	//time.Sleep(4 * time.Second)
	h.homeTemplate.Execute(w, vm)
}
