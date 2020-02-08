package main

import (
	"fmt"
	"net/http"

	"github.com/rahul-golang/crawler/handlers"
	"github.com/rahul-golang/crawler/services"
)

func main() {

	service := services.NewAppServiceImpl()
	handler := handlers.NewAppHandlers(service)

	http.HandleFunc("/", handler.ServeApplication)
	http.HandleFunc("/url", handler.GetURLCount)
	fmt.Println(http.ListenAndServe(":8000", nil))
}
