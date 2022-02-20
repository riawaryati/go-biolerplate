package infra

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	constants "github.com/furee/backend/constants/general"
	"github.com/furee/backend/domain/general"
	"github.com/furee/backend/utils"
)

type S3List struct {
	Product S3Itf
}

const (
	//Default Access duration of AWS S3 URL
	defaultURLDuration = 360 * time.Minute

	//Used to create AWS S3 Key
	encodeSecretKey = "U7aCxJLN7ykuYVJT"

	PublicAccess  = "public-read"
	PrivateAccess = "private"
)

//List of AWS S3 services that can be used in our repo
type S3ServiceItf interface {
	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
	GetObjectRequest(*s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput)
}

//List of action that will be using or needed to use AWS S3 services in our repo
type S3Itf interface {
	UploadMultiPartFile(userID int64, file *multipart.File, fileHeader *multipart.FileHeader) (string, error)
	GetURL(URLType string, key string) (string, error)
}

type S3 struct {
	region      string
	fileSource  string
	tempFolder  string
	service     S3ServiceItf
	bucketName  string
	urlDuration time.Duration
	access      string
}

func NewS3Client(cfg general.S3CredentialKey, bucketName, fileSource, access string) (S3Itf, error) {
	creds := credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, "")
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}

	awsCfg := aws.NewConfig().WithRegion(cfg.Region).WithCredentials(creds)
	svc := s3.New(session.New(), awsCfg)

	URLDuration := defaultURLDuration
	if cfg.URLDuration != 0 {
		URLDuration = time.Duration(cfg.URLDuration) * time.Minute
	}

	return S3{
		region:      cfg.Region,
		fileSource:  fileSource,
		tempFolder:  cfg.TempFolder,
		service:     svc,
		bucketName:  cfg.BucketName,
		urlDuration: URLDuration,
		access:      access,
	}, nil
}

//Will Upload file to AWS S3 & will return key + error
func (dep S3) UploadMultiPartFile(userID int64, file *multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	var key string

	filePath, err := dep.saveFiletoTempFolder(file, fileHeader)
	if err != nil {
		return key, err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return key, err
	}
	defer f.Close()

	key = dep.makeS3FilePath(userID, fileHeader.Filename)

	input := &s3.PutObjectInput{
		Bucket: aws.String(dep.bucketName),
		Key:    aws.String(key),
		Body:   f,
	}

	if dep.access == PublicAccess {
		input.ACL = aws.String(dep.access)
	}

	_, err = dep.service.PutObject(input)
	if err != nil {
		return key, err
	}

	//Remove temporary file
	err = os.Remove(filePath)
	if err != nil {
		return key, err
	}

	return key, nil
}

//Write multipart file into temporary folder
func (dep S3) saveFiletoTempFolder(file *multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	var filePath string

	data, err := ioutil.ReadAll(*file)
	if err != nil {
		return filePath, err
	}

	isExist, err := utils.DirExists(dep.tempFolder)
	if err != nil {
		return filePath, err
	}

	if !isExist {
		err = os.MkdirAll(dep.tempFolder, os.ModePerm)
		if err != nil {
			return filePath, err
		}
	}

	filePath = fmt.Sprintf("%s%s", dep.tempFolder, fileHeader.Filename)

	return filePath, ioutil.WriteFile(filePath, data, 0666)
}

//AWS S3 object data is connected through KEY
func (dep S3) makeS3FilePath(userID int64, fileName string) string {
	//Give unique constraint based on fileSource
	//To have better understanding which service responsible for an object
	key := fmt.Sprintf("%s/%s/%s", "moutar", dep.fileSource, fileName)

	return key
}

func (dep S3) GetURL(URLType string, key string) (string, error) {
	if URLType == constants.URLPublic {
		return dep.getURLPublic(key), nil
	}

	return dep.getURLLimited(key)
}

func (dep S3) getURLPublic(path string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", dep.bucketName, dep.region, path)
}

func (dep S3) getURLLimited(key string) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(dep.bucketName),
		Key:    aws.String(key),
	}
	req, _ := dep.service.GetObjectRequest(input)

	return req.Presign(dep.urlDuration)
}
