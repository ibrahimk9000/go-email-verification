package main

import (
	"fmt"
)

func verifcationLink(mail, host string) error {

	messageinfo := mailmessage{
		from:    from,
		subject: subject,
		to:      []string{mail},
	}

	token, err := createToken(mail)
	if err != nil {
		return err
	}

	link := fmt.Sprintf("%s/%s/%s", host, "verification", token)

	err = Send(link, messageinfo)
	if err != nil {
		return err
	}
	return nil
}
