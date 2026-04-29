package route

import (
	"leita/src/handler"
	"leita/src/route/problem"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func RegisterRoutes(app *fiber.App) error {
	handler, err := handler.NewHandler()
	if err != nil {
		log.Error(err)
		return err
	}

	api := app.Group("/api")

	problem.Register(api, handler.ProblemHandler)

	return nil
}
