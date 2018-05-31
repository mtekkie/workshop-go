package routing

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
)

// BaseRouter returns only business valuable routes
func BaseRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler("something")).Methods(http.MethodGet, http.MethodPost)



	return r
}

func homeHandler(s string)  func (w http.ResponseWriter, r *http.Request){
	return func (w http.ResponseWriter, r *http.Request) {
		log.Printf("%s Processing request: %s", s, r.URL.Path)
		fmt.Fprintf(w, "YEeea %s", s)

	}
}
