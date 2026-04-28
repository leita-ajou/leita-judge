package problem

import (
	"leita/src/handlers/problem"

	"github.com/gofiber/fiber/v2"
)

func RegisterProblemRoutes(api fiber.Router, handler *problem.ProblemHandler) {
	problemGroup := api.Group("/problem")
	problemGroup.Post("/submit/:problemId", handler.SubmitProblem())
	problemGroup.Post("/run/:problemId", handler.RunProblem())
}
