package middleware

import (
	"fmt"
	"os"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {

	authorizationHeader := c.Get("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	if token == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	if !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "malformed token",
		})
	}
	provider, err := oidc.NewProvider(c.Context(), os.Getenv("KEYCLOCK"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error to connect to the provider",
		})
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: os.Getenv("CLIENTID")})
	_, err = verifier.Verify(c.Context(), token)
	if err != nil {
		fmt.Println(err.Error())
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
