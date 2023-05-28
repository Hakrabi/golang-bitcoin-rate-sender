package services

import (
	"crypto/tls"
	"errors"
	"example/web-service-gin/repository"
	"strconv"

	gomail "gopkg.in/mail.v2"
)

func SendEmails() error {
	bitcoinPrice, err := GetBitcoinPrice()
	if err != nil {
		return err
	}

	emails, err := repository.GetEmails()
	if err != nil {
		return errors.New("send emails error: " + err.Error())
	}

	for _, email := range emails {
		m := gomail.NewMessage()
		m.SetHeader("From", "golang-bitcoin-rate-sender@ukr.net")
		m.SetHeader("To", email)
		m.SetHeader("Subject", "Current Bitcoin Price")
		m.SetBody("text/plain", "Current Bitcoin Price: "+strconv.FormatFloat(bitcoinPrice, 'f', 3, 64))

		d := gomail.NewDialer("smtp.ukr.net", 465, "golang-bitcoin-rate-sender@ukr.net", "tNvOy8iW2LdQdbzt")

		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		if err := d.DialAndSend(m); err != nil {
			return errors.New("send emails error: " + err.Error())
		}
	}

	return nil
}
