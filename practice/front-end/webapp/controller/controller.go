package controller

import (
	"html/template"
	"net/http"
	"strings"
)

var (
	homeController home
	shopController shop
)

// Startup : initialises handlers and registers routes
func Startup(templates map[string]*template.Template) {

	homeController.homeTemplate = templates["home.html"]
	homeController.standLocatorTemplate = templates["stand_locator.html"]
	shopController.shopTemplate = templates["shop.html"]
	shopController.categoryTemplate = templates["shop_details.html"]
	shopController.productTemplate = templates["shop_detail.html"]
	homeController.registerRoutes()
	shopController.registerRoutes()

	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))

	// Redirect /html links
	http.HandleFunc("/html/", handleOldLinks)
}

func handleOldLinks(response http.ResponseWriter, request *http.Request) {
	newURL := "http://" + request.Host + request.URL.Path[len("/html"):]
	if strings.HasSuffix(newURL, ".html") {
		newURL = newURL[:len(newURL)-5]
	}
	http.Redirect(response, request, newURL, http.StatusTemporaryRedirect)
}
