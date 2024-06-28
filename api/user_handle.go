package api

import (
	"github.com/farshidboroomand/hotel_reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "farshid khan",
		LastName:  "Boroomand khan",
	}

	return c.JSON(u)
}
