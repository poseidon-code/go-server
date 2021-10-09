package server

import (
	"log"
	"net/http"
	"os"
)

// Logger for server
var Log = log.New(os.Stdout, "go-server: ", log.LstdFlags | log.Lshortfile)


// Request Headers for CORS 
func Headers(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Content-Type", "application/json")
}


// Server settings
func New(mux *http.ServeMux, port string) *http.Server {
    return &http.Server{
        Addr: port,
        Handler: mux,
    }
}