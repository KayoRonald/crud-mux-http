package main

import (
	"net/http"

	"github.com/gorilla/mux"
	server "gtihub.com/KayoRonald/crud-mux-http/cmd"
	"gtihub.com/KayoRonald/crud-mux-http/database"
	"gtihub.com/KayoRonald/crud-mux-http/handle"
	"gtihub.com/KayoRonald/crud-mux-http/middleware"
	"gtihub.com/KayoRonald/crud-mux-http/middleware/logger"
)

func main() {
	r := mux.NewRouter()
	r.Use(logger.New)
	r.Use(middleware.CorsMiddleware)
	r.NotFoundHandler = http.HandlerFunc(handle.NotFound)
	// connect database
	database.ConnectDB()
  err := server.Run(r)
	if err != nil {
    panic(err)
  }
}