package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/farshidboroomand/hotel_reservation/api"
	"github.com/farshidboroomand/hotel_reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"
const dbName = "hotel_reservation"
const collectionName = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	collection := client.Database(dbName).Collection(collectionName)
	ctx := context.Background()
	user := types.User{
		FirstName: "farshid",
		LastName:  "borooamnd",
	}

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	listenAddr := flag.String("listenAddr", ":5000", "The Listen Address Of Api Server")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleGetUser)
	app.Listen(*listenAddr)
}
