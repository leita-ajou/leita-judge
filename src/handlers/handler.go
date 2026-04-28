package handlers

import (
	"leita/src/handlers/problem"
	"leita/src/services"

	"github.com/gofiber/fiber/v2/log"
)

type Handler struct {
	ProblemHandler *problem.ProblemHandler
}

func NewHandler() (*Handler, error) {
	service, err := services.NewService()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	problemHandler := problem.NewProblemHandler(service.ProblemService)

	return &Handler{
		ProblemHandler: problemHandler,
	}, nil
}
