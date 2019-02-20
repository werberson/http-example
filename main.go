package main

import (
	"github.com/gorilla/mux"
	"github.com/werberson/http-example/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/new/{name}", handler.KeyHandler)
	router.HandleFunc("/message/{name}", handler.MessageHandler)
	router.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
