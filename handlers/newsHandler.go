package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewsHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rw.WriteHeader(http.StatusOK)
	fmt.Println("Hello news vars: %v", vars)
}
