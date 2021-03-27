package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

const (
	_from    = "test@gmail.com"
	_subject = "test my app"
)

const (
	_smtpadr  = "smtp.gmail.com:587"
	_smtphost = "smtp.gmail.com"
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
		host:     _smtphost,
	}

	err := SendMail(sa, _smtpadr, link, ms)

	if err != nil {
		return err
	}

	return nil
}
