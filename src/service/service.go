package service

import (
	"leita/src/executor"
	"leita/src/repository"
	"leita/src/service/problem"

	"github.com/gofiber/fiber/v2/log"
)

type Service struct {
	ProblemService *problem.Service
}

func NewService() (*Service, error) {
	repository, err := repository.NewRepository()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	exec := executor.NewOsExecutor()
	problemService := problem.NewService(repository.ProblemRepository, repository.FileRepository, exec)

	return &Service{
		ProblemService: problemService,
	}, nil
}
