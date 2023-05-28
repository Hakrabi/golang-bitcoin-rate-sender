package repository

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

const repositoryFilePath = "emails.dat"

func AddEmail(email string) error {
	file, err := os.OpenFile(repositoryFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("write file error: " + err.Error())
	}
	defer file.Close()

	_, err = file.WriteString(email + "\n")
	if err != nil {
		return errors.New("write file error: " + err.Error())
	}
	return nil
}

func IsEmailExist(email string) (bool, error) {
	existingEmails, err := GetEmails()
	if err != nil {
		return false, err
	}

	for _, existingEmail := range existingEmails {
		if existingEmail == email {
			return true, nil
		}
	}
	return false, nil
}

func GetEmails() ([]string, error) {
	var emails []string
	file, err := os.OpenFile(repositoryFilePath, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return []string{}, errors.New("open file error: " + err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return []string{}, errors.New("read file error: " + err.Error())
			}
		}
		emails = append(emails, strings.TrimRight(line, "\n"))
	}
	return emails, nil
}
