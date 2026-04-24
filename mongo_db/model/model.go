package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movies struct {
	ID        primitive.ObjectId `json:"_id ,omitempty" bson:"_id ,omitempty" `
	MovieName string
	Watched   bool
}

func main() {

}
