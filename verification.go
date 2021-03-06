package main

import (
	"fmt"
)

const (
	_verification = "verification"
)

func verifcationLink(mail, host string) error {

	messageinfo := mailmessage{
		from:    _from,
		subject: _subject,
		to:      []string{mail},
	}

	token, err := createToken(mail)
	if err != nil {
		return err
	}

	link := fmt.Sprintf("%s/%s/%s", host, _verification, token)

	err = Send(link, messageinfo)
	if err != nil {
		return err
	}
	return nil
}
