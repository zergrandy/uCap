package smtp

import (
	"errors"
	"net/smtp"

	"github.com/Ranco-dev/gbms/pkg/config"
)

func Semdmail(to, title, body string) error {
	var conf = config.GetConfig()
	host := conf.GetString("smtp.hosts")
	port := conf.GetString("smtp.port")
	from := conf.GetString("smtp.from")
	pass := conf.GetString("smtp.pass")

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, pass, host)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + title + "\n" +
		body + "\n"
	err := smtp.SendMail(host+":"+port, auth, "sender@example.org", []string{to}, []byte(msg))

	if err != nil {
		return errors.New("failed to send mail: " + err.Error())
	}

	return nil
}
