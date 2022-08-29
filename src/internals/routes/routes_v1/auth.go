package routesv1

import (
	"github.com/Ashkan4472/google_form_go/src/internals/models"
	"github.com/Ashkan4472/google_form_go/src/internals/utils"
	"github.com/Ashkan4472/google_form_go/src/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	r := router.Group("/auth")
	r.Post("/signup", signUp)
	// TODO: add login
}

func signUp(c *fiber.Ctx) error {
	type UserInput struct {
		Email	string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=3,max=24"`
		FirstName string `json:"firstName" validate:"required"`
		LastName string `json:"lastName" validate:"required"`
	}

	var userInput UserInput
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad body input",
			"error": err.Error(),
		})
	}

	if err := utils.ValidateStruct(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad body input",
			"error": err,
		})
	}

	user := models.User{
		FirstName: userInput.FirstName,
		LastName: userInput.LastName,
		Email: userInput.Email,
		Password: userInput.Password,
	}
	user.HashPassword()

	tx := config.DB.Create(&user)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Transaction has been failed",
			"error": tx.Error.Error(),
		})
	}

	token, err := utils.JWTGenerate(user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
		"user": user,
	})
}
