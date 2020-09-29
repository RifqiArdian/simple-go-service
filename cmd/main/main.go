package main

import (
	"log"
	"net/http"

	"github.com/rifqiardian/simple-go-service/pkg/routes"
	"github.com/gorilla/mux"
	_"github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
