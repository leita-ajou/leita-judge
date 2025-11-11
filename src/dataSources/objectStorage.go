package dataSources

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
)

type ObjectStorage struct {
	Client objectstorage.ObjectStorageClient
}

func NewObjectStorage() (*ObjectStorage, error) {
	config := common.DefaultConfigProvider()
	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(config)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &ObjectStorage{
		Client: client,
	}, nil
}

func (objStorage *ObjectStorage) GetObject(objectName string) ([]byte, error) {
	request := objectstorage.GetObjectRequest{
		NamespaceName: common.String(os.Getenv("OS_NAMESPACE")),
		BucketName:    common.String(os.Getenv("OS_BUCKET")),
		ObjectName:    common.String(objectName),
	}

	response, err := objStorage.Client.GetObject(context.Background(), request)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	content, err := io.ReadAll(response.Content)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return content, nil
}

func (objStorage *ObjectStorage) PutObject(objectName string, data []byte) error {
	request := objectstorage.PutObjectRequest{
		NamespaceName: common.String(os.Getenv("OS_NAMESPACE")),
		BucketName:    common.String(os.Getenv("OS_BUCKET")),
		ObjectName:    common.String(objectName),
		PutObjectBody: io.NopCloser(bytes.NewReader(data)),
		ContentType:   common.String("text/plain"),
	}

	_, err := objStorage.Client.PutObject(context.Background(), request)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (objStorage *ObjectStorage) ListObjects(folderPath string) ([]objectstorage.ObjectSummary, error) {
	request := objectstorage.ListObjectsRequest{
		NamespaceName: common.String(os.Getenv("OS_NAMESPACE")),
		BucketName:    common.String(os.Getenv("OS_BUCKET")),
		Prefix:        common.String(folderPath),
	}

	response, err := objStorage.Client.ListObjects(context.Background(), request)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return response.ListObjects.Objects, nil
}
