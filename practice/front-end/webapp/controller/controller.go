package controller

import (
	"html/template"
	"net/http"
	"strings"
)

var (
	homeController     home
	loginController    login
	shopController     shop
	standMapController standLocator
)

// Startup : initialises handlers and registers routes
func Startup(templates map[string]*template.Template) {

	homeController.homeTemplate = templates["home.html"]
	homeController.registerRoutes()

	loginController.loginTemplate = templates["login.html"]
	loginController.registerRoutes()

	shopController.shopTemplate = templates["shop.html"]
	shopController.categoryTemplate = templates["shop_details.html"]
	shopController.productTemplate = templates["shop_detail.html"]
	shopController.registerRoutes()

	standMapController.standMapTemplate = templates["stand_locator.html"]
	standMapController.registerRoutes()

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
