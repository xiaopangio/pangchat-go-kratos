package pkg

import (
	"crypto/rand"
	"fmt"
	"github.com/lithammer/shortuuid/v3"
)

const (
	LettersLetter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LettersNumber = "0123456789"
	LettersSymbol = "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
)

func RandString(n int, letters ...string) (string, error) {

	lettersDefaultValue := LettersLetter + LettersNumber + LettersSymbol

	if len(letters) > 0 {
		lettersDefaultValue = ""
		for _, letter := range letters {
			lettersDefaultValue = lettersDefaultValue + letter
		}
	}

	bytes := make([]byte, n)

	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = lettersDefaultValue[b%byte(len(lettersDefaultValue))]
	}

	return string(bytes), nil
}
func UuidShort() string {
	return shortuuid.New()
}

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

func RandInt(n int, letters ...string) (string, error) {
	lettersDefaultValue := LettersNumber
	if len(letters) > 0 {
		lettersDefaultValue = ""
		for _, letter := range letters {
			lettersDefaultValue = lettersDefaultValue + letter
		}
	}
	bytes := make([]byte, n)

	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = lettersDefaultValue[b%byte(len(lettersDefaultValue))]
	}

	return string(bytes), nil
}
func Code() (string, error) {
	code, err := RandInt(4)
	if err != nil {
		return "", fmt.Errorf("cannot generate sms code: %w", err)
	}
	return code, nil
}
