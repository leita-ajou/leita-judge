package problem

import (
	"leita/src/handler/problem"

	"github.com/gofiber/fiber/v3"
)

func Register(api fiber.Router, handler *problem.Handler) {
	problemGroup := api.Group("/problem")
	problemGroup.Post("/submit/:problemId", handler.SubmitProblem())
	problemGroup.Post("/run/:problemId", handler.RunProblem())
}
