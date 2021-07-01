package main

import (
	"log"
	"net/http"

	"github.com/Metalscreame/go-training/day_6/networking-handlers/internal/middlware"
	"github.com/Metalscreame/go-training/day_6/networking-handlers/internal/repository/inmemory"
	"github.com/Metalscreame/go-training/day_6/networking-handlers/internal/server"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Main function
func main() {
	logger, _ := zap.NewProduction()
	repo := inmemory.NewRepository()
	srv := server.NewServer(repo, logger)
	md := &middlware.Middleware{Logger: logger}
	// Init router
	r := mux.NewRouter()

	// type Handler interface {
	//    ServeHTTP(ResponseWriter, *Request)
	//}
	//http.HandleFunc("/", h1)
	//	http.HandleFunc("/endpoint", h2)
	//https://golang.org/pkg/net/http/#HandleFunc

	// we can also use middleware
	r.Use(md.LoggingMiddleware)

	// Route handles & endpoints
	r.HandleFunc("/books", srv.GetBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", srv.GetBook).Methods("GET")
	r.HandleFunc("/books", srv.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", srv.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", srv.DeleteBook).Methods("DELETE")
	logger.Info("starting web server")

	// Start server
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
