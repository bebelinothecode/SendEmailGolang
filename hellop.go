package main

import (
	"log"
	"net/smtp"
	"os"
	"github.com/joho/godotenv"
)

func sendMail() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Failed to load .env values:%s", err)
	}
	
	msg := "hello from the other side"

	auth := smtp.PlainAuth(
		"",
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
	)

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"bebelino@gmail.com",
		[]string{os.Getenv("USERNAME")},
		[]byte(msg),
	)

	if err != nil {
		log.Fatalf("Error occured sending email:",err)
	}

	log.Println("Email sent!!!!")
}

func main() {
	sendMail()
}