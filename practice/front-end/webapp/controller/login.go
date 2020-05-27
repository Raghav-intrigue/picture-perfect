package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"clumio.com/roger/picture-perfect/webapp/viewmodel"
)

type login struct {
	loginTemplate *template.Template
}

func (l login) registerRoutes() {
	http.HandleFunc("/login", l.handleLogin)
}

func (l login) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error logging in: %v", err))
		}

		email := r.Form.Get("email")
		password := r.Form.Get("password")

		if email == "test@gmail.com" && password == "password" {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		}

		vm.Email = email
		vm.Password = password
	}
	w.Header().Add("Content-Type", "text/html")
	l.loginTemplate.Execute(w, vm)
}
