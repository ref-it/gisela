package common

import (
	"fmt"
	"net/smtp"
	"strconv"
)

func SendEmail(body string, to []string) {
	config := GetSmtpConfiguration()

	// Authentication.
	auth := smtp.PlainAuth("", config.User, config.Password, config.Host)

	// Sending email.
	err := smtp.SendMail(config.Host+":"+strconv.Itoa(config.Port), auth, config.SenderAddress, to, []byte(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
