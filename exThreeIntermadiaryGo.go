package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	from := "sender@email.com"
	password := "yourPassword"
	to := "recipient@email.com"
	subject := "Test"
	message := "Hello, Go!"

	msg := []byte(fmt.Sprintf("Subject: %s\n\n%s\n", subject, message))

	err := smtp.SendMail("smtp.example.com:587",
		auth(from, password),
		from,
		[]string{to},
		msg,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email sucessfully sent!")
}

func auth(username, password string) smtp.Auth {
	return smtp.PlainAuth("", username, password, "")
}
