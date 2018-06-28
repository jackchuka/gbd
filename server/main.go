package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/jackchuka/gbd/configs"
	"github.com/jackchuka/gbd/handler"
)

func main() {
	r := mux.NewRouter()

	h := handler.Handler{
		Configs: configs.GetConfigs(),
	}

	r.HandleFunc("/api", h.APIHandler).Methods("POST")

	// server web static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("web/dist")))

	// Bind to a port and pass our router in
	log.Println("server listining...")
	log.Println(http.ListenAndServe(":8888", r))
}
