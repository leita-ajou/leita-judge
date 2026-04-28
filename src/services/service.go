package services

import (
	"leita/src/repositories"
	"leita/src/services/problem"

	"github.com/gofiber/fiber/v2/log"
)

type Service struct {
	ProblemService *problem.ProblemService
}

func NewService() (*Service, error) {
	repository, err := repositories.NewRepository()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	problemService := problem.NewProblemService(repository.ProblemRepository)

	return &Service{
		ProblemService: problemService,
	}, nil
}
