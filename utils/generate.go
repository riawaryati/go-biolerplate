package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"regexp"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUUID() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}

func GeneratePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(savedPass, incomingPass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(savedPass), []byte(incomingPass))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func NewInvoice(userID int64, pattern string) string {
	timeNow := time.Now().UTC()
	month := ConvertMonthtoRoman(int(timeNow.Month()))

	invoice := fmt.Sprintf("%s%v/%s/%v", pattern, userID, month, timeNow.Unix())

	return invoice
}

func GetExtFilename(filename string) (string, error) {
	r, err := regexp.Compile(`\.([a-z0-9]+)$`)
	if err != nil {
		return "", err
	}

	imgext := r.FindString(filename)
	return imgext, nil
}

func GenerateOTP() string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, 6)
	n, err := io.ReadAtLeast(rand.Reader, b, 6)
	if n != 6 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
