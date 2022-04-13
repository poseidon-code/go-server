package server

import (
	"os"

	dotenv "github.com/joho/godotenv"
)

// Environment Variables structure for reading environment variables
// PORT (required) : eg. PORT=5000
// TODO : add new entries for every new environment variables in '.env'
type EnvironmentVariables struct {
    // (Heroku deployment of any Go based REST server does not requires to manually set a 'PORT' varaible in Heroku's config,
    // instead it is mandatory to handle 'PORT' environment variable provided by Heroku, inside the program, using os.Getenv("PORT"))
    // DO NOT REMOVE 'Port'
    Port    string

    // other .env fields should also be added for convenience of accessing/structuring them
    MDBUserName     string
    MDBPassword     string
}


var ENV *EnvironmentVariables


// Loads Environment variables before working with server.
// Requires default 'PORT' to be passed in
func LoadEnvironment(port string) {
    // reading `.env` config file in root project directory
    err := dotenv.Load(); if err!=nil {
        Log.Println("Error loading `.env` config file :\n", err, "\n(Ignore this error if there are no specific environment variables set/required for this server)")
        Log.Println("Starting without environment variables")
    }


    // setting environment variables
    ENV = &EnvironmentVariables {
        // Checks & sets, if 'PORT' environment variable is set, else sets passed 'PORT' value
        // this checks enables to comply with Heroku 'PORT' configuration when deploying to Heroku,
        // even if during development one could set any value for 'PORT',
        // but in production, that 'PORT' value will be set by Heroku using environment variables,
        // thats why it is mandatory to handle the provided 'PORT' value by Heroku.

        // setting 'PORT' from environment varables,
        // if no 'PORT' found in environment, then set passed 'PORT' value
        // NO NEED TO EDIT THIS
        Port: func() string { if value, ok := os.LookupEnv("PORT"); ok {return value} else {return port} }(),

        MDBUserName: os.Getenv("MONGODB_USERNAME"),
        MDBPassword: os.Getenv("MONGODB_PASSWORD"),
    }
}