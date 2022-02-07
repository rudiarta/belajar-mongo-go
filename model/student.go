package model

import "time"

type Student struct {
	Name      string    `bson:"name"`
	Age       int       `bson:"age"`
	Address   string    `bson:"address"`
	CreatedAt time.Time `bson:"created_at"`
}
