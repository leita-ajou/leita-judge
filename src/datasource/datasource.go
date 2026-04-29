package datasource

import (
	"github.com/gofiber/fiber/v2/log"
)

type DataSource struct {
	objectStorage *ObjectStorage
}

func NewDataSource() (*DataSource, error) {
	os, err := NewObjectStorage()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &DataSource{
		objectStorage: os,
	}, nil
}

func (ds *DataSource) GetObjectStorage() *ObjectStorage {
	return ds.objectStorage
}
