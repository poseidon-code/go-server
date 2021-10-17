package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	s "github.com/poseidon-code/go-server/server"
)

var CTX = context.TODO()

// MongoDB URI string (required)
// the mongoDB connection URI string, needs to be set before accessing database,
// (preferably before creating server)
var URI string

// MongoDB client (globally used)
// globally accesed while doing operations with the database,
// after connecting (Connect()), set by Connect()
var Client *mongo.Client


func Connect() {
    client, err := mongo.NewClient(options.Client().ApplyURI(URI)); if err!=nil {
        s.Log.Fatalln("Error connecting to database :\n", err)
	}
    client.Connect(CTX)
    Client = client
}

func Disconnect() {
    Client.Disconnect(CTX)
}