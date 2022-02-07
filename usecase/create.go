package usecase

import (
	"context"
	"fmt"

	"github.com/rudiarta/belajar-mongo-go/connection"
	"github.com/rudiarta/belajar-mongo-go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDocument(data model.Student) (string, error) {
	conn := connection.Conn

	result, err := conn.Master().Collection("students").InsertOne(context.Background(), data)
	if err != nil {
		return "", err
	}

	res := result.InsertedID.(primitive.ObjectID)
	fmt.Printf("ObjectID: %s \n", res.Hex())

	return res.Hex(), nil
}
