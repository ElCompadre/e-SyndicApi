package main

import (
	"fmt"
	"log"
	"net/http"

	"./controllers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("e-Syndic V0.0")
	router := mux.NewRouter()

	router.HandleFunc("/copros", controllers.GetCoproByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+conf.Postgre.SitePort, router))
}
