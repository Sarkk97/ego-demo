package main

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

//Mailer is an email sender
type Mailer interface {
	Send()
}

type smtpMailer struct {
	from    string
	to      []string
	bcc     []string
	cc      []string
	subject string
	body    []byte
}

//Send sends an mail message using SMTP
func (mailer *smtpMailer) Send() {
	e := email.NewEmail()
	e.From = mailer.from
	e.To = mailer.to
	e.Bcc = mailer.bcc
	e.Cc = mailer.cc
	e.Subject = mailer.subject
	e.HTML = mailer.body

	/*var server, port, emailID, password string

	var exist bool

	if server, exist = os.LookupEnv("SMTP_SERVER"); exist == false {
		log.Panicln("SMTP_SERVER env variable not set")
	}

	if port, exist = os.LookupEnv("SMTP_PORT"); exist == false {
		log.Panicln("SMTP_PORT env variable not set")
	}

	if emailID, exist = os.LookupEnv("EMAIL_ID"); exist == false {
		log.Panicln("EMAIL_ID env variable not set")
	}

	if password, exist = os.LookupEnv("EMAIL_PASSWORD"); exist == false {
		log.Panicln("EMAIL_PASSWORD env variable not set")
	}*/

	//err := e.Send(fmt.Sprintf("%s:%s", server, port), smtp.PlainAuth("", emailID, password, server))
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "devs.ego@gmail.com", "gopher220", "smtp.gmail.com"))

	if err != nil {
		log.Println(err.Error())
	}
}

//NewMailer returns an instance of a Mailer implementation
func NewMailer(
	from string,
	to []string,
	bcc []string,
	cc []string,
	subject string,
	body []byte,
) Mailer {
	return &smtpMailer{
		from:    from,
		to:      to,
		bcc:     bcc,
		cc:      cc,
		subject: subject,
		body:    body,
	}
}
