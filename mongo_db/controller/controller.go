package controller

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connString = "mongodb+srv://devprajwolgotnochill:<8MHyR9EoAZBFS287>@letsgomongo.yc0ie7k.mongodb.net/?appName=letsgomongo"
const dbName = "Movies"
const colname = "Watchlist"

//imp

var controller *mongo.Collection

func init() {
	//client syntax
	opts := options.Client().ApplyURI(connString)

	//
	client, err := mongo.Connect(opts)
}
