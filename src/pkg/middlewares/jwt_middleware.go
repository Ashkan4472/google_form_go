package middlewares

import (
	"os"

	"github.com/Ashkan4472/google_form_go/src/internals/models"
	"github.com/Ashkan4472/google_form_go/src/pkg/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

var JwtMiddleWare = func () func(*fiber.Ctx) error {
	jwtSecret := os.Getenv("JWT_USER_SECRET")

	return jwtware.New(jwtware.Config{
			SigningKey: []byte(jwtSecret),
			SuccessHandler: func (c *fiber.Ctx) error {
				userLocal := c.Locals("user").(*jwt.Token)
				claims := userLocal.Claims.(jwt.MapClaims)
				var user models.User
				config.DB.Find(&user, "id = ? AND email = ?", claims["id"], claims["email"])
				c.Locals("user", user)
				return c.Next()
			},
		})
}
