package main

import (
	"time"

	"github.com/rudiarta/belajar-mongo-go/model"
	"github.com/rudiarta/belajar-mongo-go/usecase"
)

func main() {
	id, _ := usecase.CreateDocument(model.Student{
		Name:      "Murid Baru lagi",
		Age:       12,
		Address:   "Jalan menuju yang benar",
		CreatedAt: time.Now(),
	})
	usecase.ReadOneDocument(id)

}
