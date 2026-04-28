package routes

import (
	"leita/src/handlers"
	"leita/src/routes/problem"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func RegisterRoutes(app *fiber.App) error {
	handler, err := handlers.NewHandler()
	if err != nil {
		log.Error(err)
		return err
	}

	api := app.Group("/api")

	problem.RegisterProblemRoutes(api, handler.ProblemHandler)

	return nil
}
