package main

import (
	"cash/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {

	routersInit := routers.InitRouter()
	maxHeaderBytes := 1 << 20
	endPoint := fmt.Sprintf(":%d", 8000)
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    60,
		WriteTimeout:   60,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
