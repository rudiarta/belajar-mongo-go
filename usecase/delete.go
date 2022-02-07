package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/rudiarta/belajar-mongo-go/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Delete(objectID string) error {
	conn := connection.Conn

	id, _ := primitive.ObjectIDFromHex(objectID)
	filter := bson.D{{"_id", id}}
	result, err := conn.Master().Collection("students").DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

	return nil
}
