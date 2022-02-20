package utils

import (
	"mime/multipart"
	"net/http"
	"os"
	"regexp"

	domain "github.com/furee/backend/domain/general"
)

func PhoneNumberValidator(phone string) bool {
	val := regexp.MustCompile(`[^0-9]*1[34578][0-9]{9}[^0-9]*`)
	return val.MatchString(phone)
}

func ImageValidator(image multipart.File, header *multipart.FileHeader, imageSize int64) (bool, string) {
	if header.Size > imageSize {
		return false, "image too large, max size 1 MB"
	}

	// Create a buffer to store the header of the file in
	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := image.Read(fileHeader); err != nil {
		return false, "fail to get image data"
	}

	// set position back to start.
	if _, err := image.Seek(0, 0); err != nil {
		return false, "fail to get image data"
	}

	if !domain.IsAllowImageType(http.DetectContentType(fileHeader)) {
		return false, "image type invalid"
	}

	return true, ""
}

func FileValidator(file multipart.File, header *multipart.FileHeader, fileSize int64) (bool, string, string) {
	if header.Size > fileSize {
		return false, "", "file too large, max size 1 MB"
	}

	// Create a buffer to store the header of the file in
	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		return false, "", "fail to get file data"
	}

	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		return false, "", "fail to get file data"
	}

	mimeType, isAllow := domain.IsAllowMediaType(http.DetectContentType(fileHeader))
	if !isAllow {
		return false, "", "image type invalid"
	}

	return true, mimeType, ""
}

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
