package repositories

import (
	"leita/src/dataSources"

	"github.com/gofiber/fiber/v2/log"
)

type ProblemRepository struct {
	dataSource *dataSources.DataSource
}

func NewProblemRepository() (*ProblemRepository, error) {
	dataSource, err := dataSources.NewDataSource()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &ProblemRepository{
		dataSource: dataSource,
	}, nil
}



func (repository *ProblemRepository) SaveCode(path string, code []byte) error {
	os := repository.dataSource.GetObjectStorage()
	if err := os.PutObject(path, code); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (repository *ProblemRepository) GetObjectsInFolder(folderPath string) ([][]byte, error) {
	os := repository.dataSource.GetObjectStorage()
	objects, err := os.ListObjects(folderPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	contents := make([][]byte, 0, len(objects))
	for _, object := range objects {
		content, err := os.GetObject(*object.Name)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		contents = append(contents, content)
	}

	return contents, nil
}
