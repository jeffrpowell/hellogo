package handlers

import (
	"log"
	"net/http"

	"github.com/jeffrpowell/hellogo/internal/constants"
	"github.com/jeffrpowell/hellogo/internal/database"
	"github.com/jeffrpowell/hellogo/internal/handlers/helper"
	"github.com/jeffrpowell/hellogo/internal/handlers/middleware"
	"github.com/jeffrpowell/hellogo/web"
)

func init() {
	constants.ROUTER.HandleFunc("/list", middleware.DefaultMiddlewareChain(listsGET)).Methods("GET")
	constants.ROUTER.HandleFunc("/list/{listId:[0-9]+}", middleware.Chain(listHandler, append(middleware.DefaultMiddlewareSlice, middleware.ListIdOwner("listId"))...))
	constants.ROUTER.HandleFunc("/list/{listId:[0-9]+}/edit", middleware.Chain(editListGET, append(middleware.DefaultMiddlewareSlice, middleware.ListIdOwner("listId"))...)).Methods("GET")
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		listPOST(w, r)
	case "DELETE":
		listDELETE(w, r)
	case "GET":
		listGET(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

/* Get all lists of a user */
func listsGET(w http.ResponseWriter, r *http.Request) {
	userId, err := helper.GetUserId(r)
	if err != nil {
		http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	lists, err := database.GetLists(userId)
	if err != nil {
		http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	admin := helper.IsUserAdmin(r)
	listsParams := web.ListsPageParams(lists, admin)
	web.ListsPage(w, listsParams)
}

/* Edit list page */
func editListGET(w http.ResponseWriter, r *http.Request) {
	listId, _ := helper.GetPathVarInt(r, "listId") //err will trip in listIdOwner middleware first
	list, err := database.GetList(listId)
	if err != nil {
		http.Error(w, "Unexpected error occurred", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	admin := helper.IsUserAdmin(r)
	editListPageParams := web.EditListParams(list, admin)
	web.EditListPage(w, editListPageParams)
}

// the following code is stubbed to allow for the example handler code to build
func listGET(w http.ResponseWriter, r *http.Request)    {}
func listPOST(w http.ResponseWriter, r *http.Request)   {}
func listDELETE(w http.ResponseWriter, r *http.Request) {}
