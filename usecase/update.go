package usecase

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/rudiarta/belajar-mongo-go/connection"
	"github.com/rudiarta/belajar-mongo-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Update(objectID string, data model.Student) error {
	conn := connection.Conn

	id, _ := primitive.ObjectIDFromHex(objectID)
	filter := bson.D{{"_id", id}}

	update := bson.D{}
	if !strings.EqualFold(data.Name, ``) {
		update = append(update, bson.E{"$set", bson.D{{"name", data.Name}}})
	}
	if !strings.EqualFold(data.Address, ``) {
		update = append(update, bson.E{"$set", bson.D{{"address", data.Address}}})
	}
	if data.Age > 0 {
		update = append(update, bson.E{"$set", bson.D{{"age", data.Age}}})
	}
	if !data.CreatedAt.IsZero() {
		update = append(update, bson.E{"$set", bson.D{{"created_at", data.CreatedAt}}})
	}

	result, err := conn.Master().Collection("students").UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("UpdateOne update %v document(s)\n", result.MatchedCount)
	return nil
}
