package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

const (
	from    = "test@gmail.com"
	subject = "test my app"
)

const (
	smtpAdrr = "smtp.gmail.com:587"
	smtpHost = "smtp.gmail.com"
)

type mailmessage struct {
	from    string
	subject string
	to      []string
}
type smtpAuth struct {
	identity string
	username string
	password string
	host     string
}

func SendMail(sa smtpAuth, smtpAdrr, link string, ms mailmessage) error {
	e := email.NewEmail()
	e.From = ms.from
	e.To = ms.to

	e.Subject = ms.subject

	hre := fmt.Sprintf(`<a href="%s/">%s/</a>`, link, link)
	e.HTML = []byte(hre)
	err := e.Send(smtpAdrr, smtp.PlainAuth(sa.identity, sa.username, sa.password, sa.host))
	return err

}

func Send(link string, ms mailmessage) error {

	emailUser := os.Getenv("SMTPUSER")
	emailPass := os.Getenv("SMTPPASS")

	sa := smtpAuth{
		identity: "",
		username: emailUser,
		password: emailPass,
		host:     smtpHost,
	}

	err := SendMail(sa, smtpAdrr, link, ms)

	if err != nil {
		return err
	}

	return nil
}
