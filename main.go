package main

import (
	"fmt"
	"net/http"

	"./src"
)

func main() {
	fmt.Println("CCAC API Start")
	router := api.NewRouter()
	http.ListenAndServe(":8080", router)
}
