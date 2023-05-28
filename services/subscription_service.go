package services

import (
	"errors"
	"example/web-service-gin/repository"
	"net/mail"
)

func SubscribeEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("subscribe email error: " + err.Error())
	}

	isEmailExist, err := repository.IsEmailExist(email)
	if err != nil {
		return err
	}
	if isEmailExist {
		return errors.New("email exist")
	}

	err = repository.AddEmail(email)
	if err != nil {
		return err
	}
	return nil
}
