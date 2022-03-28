package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("server start")
	router := mux.NewRouter().StrictSlash(true)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", 3000), router); err != nil {
		return
	}
}
