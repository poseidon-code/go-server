package main

import (
	"fmt"
	"net/http"

	d "github.com/poseidon-code/go-server/database"
	r "github.com/poseidon-code/go-server/routes"
	s "github.com/poseidon-code/go-server/server"
)


func main() {
    mux := http.NewServeMux()           // create Mux (router)
    server := s.New(mux, "5000")        // create http.Server with config (app)

    // set up database URI
    d.URI = fmt.Sprintf("mongodb+srv://%s:%s@stories.l6tlk.mongodb.net/stories?retryWrites=true&w=majority", s.ENV.MDBUserName, s.ENV.MDBPassword)
    

    // ROUTES
    mux.HandleFunc("/", r.Home)
    mux.HandleFunc("/user/", r.User)


    // server starting & listening
    s.Log.Println("Starting Server on http://localhost"+server.Addr)
    err := server.ListenAndServe(); if err!=nil {
        s.Log.Fatalln("Error Starting Server at http://localhost"+server.Addr+" :\n", err)
    }
}