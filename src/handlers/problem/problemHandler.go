package problem

import (
	"leita/src/services/problem"
)

type ProblemHandler struct {
	service *problem.ProblemService
}

func NewProblemHandler(service *problem.ProblemService) *ProblemHandler {
	return &ProblemHandler{
		service: service,
	}
}
