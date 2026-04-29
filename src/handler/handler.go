package handler

import (
	"leita/src/handler/problem"
	"leita/src/service"

	"github.com/gofiber/fiber/v3/log"
)

type Handler struct {
	ProblemHandler *problem.Handler
}

func NewHandler() (*Handler, error) {
	service, err := service.NewService()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	problemHandler := problem.NewHandler(service.ProblemService)

	return &Handler{
		ProblemHandler: problemHandler,
	}, nil
}
