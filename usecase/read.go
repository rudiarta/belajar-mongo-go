package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rudiarta/belajar-mongo-go/connection"
	"github.com/rudiarta/belajar-mongo-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReadOneDocument(objectID string) (model.Student, error) {
	conn := connection.Conn

	id, _ := primitive.ObjectIDFromHex(objectID)
	var result bson.M
	filter := bson.D{{"_id", id}}
	conn.Master().Collection("students").FindOne(context.Background(), filter).Decode(&result)

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	res := model.Student{}
	json.Unmarshal(output, &res)

	fmt.Println(res)

	return res, nil
}
