package middleware

import (
	"github.com/agilistikmal/safatanc-go/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
)

func ProjectMiddleware(ctx *fiber.Ctx) error {
	apiKeyRequest := string(ctx.Request().Header.Peek("API-Key"))
	if apiKeyRequest != os.Getenv("API_KEY") {
		return ctx.JSON(model.Response{
			StatusCode:    http.StatusUnauthorized,
			StatusMessage: http.StatusText(http.StatusUnauthorized),
		})
	}
	return ctx.Next()
}
