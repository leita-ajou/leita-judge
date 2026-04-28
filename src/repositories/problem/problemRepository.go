package problem

import (
	"leita/src/dataSources"

	"github.com/gofiber/fiber/v2/log"
)

type ProblemRepository struct {
	dataSource *dataSources.DataSource
}

func NewProblemRepository(dataSource *dataSources.DataSource) *ProblemRepository {
	return &ProblemRepository{
		dataSource: dataSource,
	}
}

func (repository *ProblemRepository) SaveCode(path string, code []byte) error {
	os := repository.dataSource.GetObjectStorage()
	if err := os.PutObject(path, code); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (repository *ProblemRepository) GetObjectsInFolder(path string) ([][]byte, error) {
	os := repository.dataSource.GetObjectStorage()
	objects, err := os.GetObjectsInFolder(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return objects, nil
}
