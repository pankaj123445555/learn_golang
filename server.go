package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := ":8080"
	r := mux.NewRouter()

	SetUpRouter(r)

	fmt.Println(`server is running on Port :`, port)
	log.Fatal(http.ListenAndServe(port, r))
}
