package handlers

import (
	"net/http"

	"github.com/jeffrpowell/hellogo/internal/constants"
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
	param := "name"
	if values, ok := query[param]; ok && len(values) > 0 {
		web.HelloWorldPage(w, web.HelloWorldParams(values[0]))
	} else {
		web.HelloWorldPage(w, web.HelloWorldParams("World"))
	}
}
