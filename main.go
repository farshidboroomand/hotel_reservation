package main

import (
	"flag"

	"github.com/farshidboroomand/hotel_reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The Listen Address Of Api Srver")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleGetUser)
	app.Listen(*listenAddr)
}
