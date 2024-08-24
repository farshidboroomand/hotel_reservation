package api

import (
	"errors"

	"github.com/farshidboroomand/hotel_reservation/db"
	"github.com/farshidboroomand/hotel_reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	UserStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		UserStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)
	user, err := h.UserStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "not found"})
		}
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.UserStore.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) HandleCreateNewUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}

	user, err := types.CreateNewUserFromparams(params)

	if err != nil {
		return err
	}

	insertedUser, err := h.UserStore.InsertUser(c.Context(), user)

	if err != nil {
		return nil
	}

	return c.JSON(insertedUser)
}
