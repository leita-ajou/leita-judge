package repositories

import (
	"leita/src/dataSources"
	"leita/src/repositories/problem"

	"github.com/gofiber/fiber/v2/log"
)

type Repository struct {
	ProblemRepository *problem.ProblemRepository
}

func NewRepository() (*Repository, error) {
	dataSource, err := dataSources.NewDataSource()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	repository := problem.NewProblemRepository(dataSource)

	return &Repository{
		ProblemRepository: repository,
	}, nil
}
