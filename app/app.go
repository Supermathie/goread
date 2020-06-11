package main

import (
	"net/http"

	"github.com/gorilla/mux"

	app "github.com/mblakele/goread"

	"google.golang.org/appengine"
)

func init() {
	router := mux.NewRouter()
	app.RegisterHandlers(router)
	http.Handle("/", router)
}

/* compatibility for go111 */
func main() {
	appengine.Main()
}
