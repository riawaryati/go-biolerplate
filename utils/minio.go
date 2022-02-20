// Amazon S3 Compatible Cloud Storage

package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/furee/backend/domain/general"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	PublicAccess  = "public-read"
	PrivateAccess = "private"
)

type Minio struct {
	client     *minio.Client
	bucket     string
	tempFolder string
	baseURL    string
}

func NewMinio(cfg general.MinioSecret) (Minio, error) {
	client, err := minio.New(
		cfg.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(cfg.Key, cfg.Secret, ""),
			Secure: true,
			Region: cfg.Region,
		},
	)
	if err != nil {
		return Minio{}, err
	}

	return Minio{
		client:     client,
		bucket:     cfg.BucketName,
		tempFolder: cfg.TempFolder,
		baseURL:    cfg.BaseURL,
	}, nil
}

func (m Minio) UploadMultiPartFile(access string, folderName string, file *multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	uploadPath := fmt.Sprintf("%s/%s", folderName, fileHeader.Filename)
	location := fmt.Sprintf("%s%s", m.baseURL, uploadPath)

	tempFilePath, err := m.saveFiletoTempFolder(file, fileHeader)
	if err != nil {
		return "", err
	}

	f, err := os.Open(tempFilePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	objectStat, err := f.Stat()
	if err != nil {
		return "", err
	}

	_, err = m.client.PutObject(
		context.Background(),
		m.bucket,
		uploadPath,
		f,
		objectStat.Size(),
		minio.PutObjectOptions{
			ContentType:  fileHeader.Header.Get("Content-Type"),
			UserMetadata: map[string]string{"x-amz-acl": access},
		},
	)
	if err != nil {
		return "", err
	}

	//Remove temporary file
	err = os.Remove(tempFilePath)
	if err != nil {
		return "", err
	}

	return location, nil
}

//Write multipart file into temporary folder
func (m Minio) saveFiletoTempFolder(file *multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	var filePath string

	data, err := ioutil.ReadAll(*file)
	if err != nil {
		return filePath, err
	}

	isExist, err := DirExists(m.tempFolder)
	if err != nil {
		return filePath, err
	}

	if !isExist {
		err = os.MkdirAll(m.tempFolder, os.ModePerm)
		if err != nil {
			return filePath, err
		}
	}

	filePath = fmt.Sprintf("%s%s", m.tempFolder, fileHeader.Filename)

	return filePath, ioutil.WriteFile(filePath, data, 0666)
}
