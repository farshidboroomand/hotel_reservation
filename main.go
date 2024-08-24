package main

import (
	"context"
	"log"
	"os"

	"github.com/farshidboroomand/hotel_reservation/api"
	"github.com/farshidboroomand/hotel_reservation/db"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoEndpoint := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	var (
		app         = fiber.New()
		apiV1       = app.Group("/api/v1")
		userStore   = db.NewMongoUserStore(client)
		userHandler = api.NewUserHandler(userStore)
	)

	apiV1.Get("/users/:id", userHandler.HandleGetUser)
	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Post("/users", userHandler.HandleCreateNewUser)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
