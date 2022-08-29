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
	r.Post("/login", login)
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

func login(c *fiber.Ctx) error {
	type UserInput struct {
		Email string `json:"email" validate:"required,email"`
		Password string `json:"password"`
	}

	var userInput UserInput
	if err := c.BodyParser(&userInput); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Body request",
			"error": err.Error(),
		})
	}

	if err := utils.ValidateStruct(userInput); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Body request",
			"error": err,
		})
	}

	var fetchedUser models.User
	tx := config.DB.First(&fetchedUser, "email = ?", userInput.Email)
	if tx.Error != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	isPassValid, _ := fetchedUser.CheckPasswordHash(userInput.Password)
	if !isPassValid {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token, _ := utils.JWTGenerate(fetchedUser)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"user": fetchedUser,
	})
}
