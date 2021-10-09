package main

import (
	"net/http"

	r "github.com/poseidon-code/go-server/routes"
	s "github.com/poseidon-code/go-server/server"
)


func main() {
    mux := http.NewServeMux()       // create Mux (router)
    server := s.New(mux, ":5000")   // create http.Server with config (app)


    // ROUTES
    mux.HandleFunc("/", r.Home)
    mux.HandleFunc("/user/", r.User)


    // server starting & listening
    s.Log.Println("Starting Server on http://localhost:5000")
    err := server.ListenAndServe(); if err!=nil {
    s.Log.Fatalln("Error Starting Server at http://localhost:5000 :\n", err)
    }
}