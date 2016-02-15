package email

import (
	//"fmt"
	"net/smtp"
	"strings"
)

type Email struct {
	User     string
	Pass     string
	Host     string
	To       string
	Subject  string
	Body     string
	Mailtype string
}

func (E *Email) SendToMail() error {
	hp := strings.Split(E.Host, ":")
	auth := smtp.PlainAuth("", E.User, E.Pass, hp[0])
	var content_type string
	if E.Mailtype == "html" {
		content_type = "Content-Type: text/" + E.Mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + E.To + "\r\nFrom: " + E.User + "\r\nSubject: " + E.Subject + "\r\n" + content_type + "\r\n\r\n" + E.Body)
	send_to := strings.Split(E.To, ";")
	err := smtp.SendMail(E.Host, auth, E.User, send_to, msg)
	return err
}
