package routes

import (
	"leita/src/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func RegisterProblemRoutes(api fiber.Router) error {
	handler, err := handlers.NewProblemHandler()
	if err != nil {
		log.Error(err)
		return err
	}

	problemGroup := api.Group("/problem")
	problemGroup.Post("/submit/:problemId", handler.SubmitProblem())
	problemGroup.Post("/run/:problemId", handler.RunProblem())

	return nil
}
