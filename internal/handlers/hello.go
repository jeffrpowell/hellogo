package handlers

import (
	"net/http"

	"github.com/jeffrpowell/hellogo/internal/constants"
	"github.com/jeffrpowell/hellogo/internal/database"
	"github.com/jeffrpowell/hellogo/web"
)

func init() {
	constants.ROUTER.HandleFunc("/", helloWorldGET).Methods("GET")
	constants.ROUTER.HandleFunc("/hello", helloWorldPage).Methods("GET")
}

func helloWorldGET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var name = "World"
	if values, ok := query["name"]; ok && len(values) > 0 {
		name = values[0]
	}
	var gradientName = "Foothill sunrise"
	if values, ok := query["gradient"]; ok && len(values) > 0 {
		gradientName = values[0]
	}
	gradient, err := database.GetGradient(gradientName)
	if err != nil {
		gradient = constants.ColorGradient{}
	}
	params := web.HelloWorldParams(name, gradient)
	web.HelloWorldPage(w, params)
}
