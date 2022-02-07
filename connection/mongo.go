package connection

import (
	"context"
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
	"go.elastic.co/apm/module/apmmongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Conn MongoDB

func init() {
	err := dotenv.Load(".env")
	if err != nil {
		panic(".env is not loaded properly")
	}
	Conn = InitMongoDB()
}

type mongoInstance struct {
	master *mongo.Database
}

func (c *mongoInstance) Master() *mongo.Database {
	return c.master
}
func (c *mongoInstance) Close() {
	if err := c.master.Client().Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

type MongoDB interface {
	Master() *mongo.Database
	Close()
}

func InitMongoDB() MongoDB {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    os.Getenv("MONGODB_AUTH_SOURCE"),
		Username:      os.Getenv("MONGODB_USER"),
		Password:      os.Getenv("MONGODB_PW"),
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri).SetAuth(credential).SetMonitor(apmmongo.CommandMonitor()))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	return &mongoInstance{
		master: client.Database(os.Getenv("MONGODB_DATABASE")),
	}
}
