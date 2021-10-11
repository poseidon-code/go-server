package server

import (
	"log"
	"net/http"
	"os"

	dotenv "github.com/joho/godotenv"
)

// Logger for server
var Log = log.New(os.Stdout, "go-server: ", log.LstdFlags | log.Lshortfile)


// Environment Variables structure for reading environment variables
// PORT : (required) eg. PORT=5000
// (Heroku deployment of any Go based REST server does not requires to manually set a 'PORT' varaible in Heroku's config,
// instead it is mandatory to handle 'PORT' environment variable provided by Heroku, inside the program, using os.Getenv("PORT"))
// ADD : new entries for every new environment variables in '.env'
type EnvironmentVariables struct {
    // Checks & sets, if 'PORT' environment variable is set, else sets passed 'PORT' value from server.New()
    // this checks enables to comply with Heroku 'PORT' configuration when deploying to Heroku,
    // even if during development one could set any value for 'PORT',
    // but in production, that 'PORT' value will be set by Heroku using environment varibales
    // (assuming this case, it will be set by Heroku by default)
    Port    string
}


// Request Headers for CORS 
func Headers(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Content-Type", "application/json")
}


// Server settings
func New(mux *http.ServeMux, port string) *http.Server {
    // reading `.env` config file in root project directory
    err := dotenv.Load(); if err!=nil {
        Log.Println("Error loading `.env` config file :\n", err, "\n(Ignore this error if there are no specific environment variables set/required for this server)")
        Log.Println("Starting without environment variables")
    }


    // setting environment variables
    env := EnvironmentVariables {
        // setting 'PORT' from environment varables,
        // if no 'PORT' found in environment, then set passed 'PORT' value 
        Port: func() string {
                if value, ok := os.LookupEnv("PORT"); ok {
                    return value
                } else {
                    return port
                }
            }(),
    }

    return &http.Server{
        Addr: ":"+env.Port,
        Handler: mux,
    }
}