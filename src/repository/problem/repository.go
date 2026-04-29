package problem

import (
	"leita/src/datasource"

	"github.com/gofiber/fiber/v2/log"
)

type Repository struct {
	dataSource *datasource.DataSource
}

func NewRepository(dataSource *datasource.DataSource) *Repository {
	return &Repository{
		dataSource: dataSource,
	}
}

func (repository *Repository) SaveCode(path string, code []byte) error {
	os := repository.dataSource.GetObjectStorage()
	if err := os.PutObject(path, code); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (repository *Repository) GetObjectsInFolder(path string) ([][]byte, error) {
	os := repository.dataSource.GetObjectStorage()
	objects, err := os.GetObjectsInFolder(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return objects, nil
}
