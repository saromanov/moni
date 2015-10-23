package moni

import(
  "gopkg.in/gomail.v2"
)

type Email struct {
	SmtpServer string
	FromAddress string
	ToAddress string
	Username string
	Passowrd string
}

func (email *Email) Send(message string) {
	m := gomail.NewMessage()
    m.SetAddressHeader("From", email.FromAddress, "Moni")
    m.SetAddressHeader("To", email.ToAddress, "test")
    m.SetHeader("Subject", "Notification from Moni")
    m.SetBody("text/plain", message)

    d := gomail.NewPlainDialer(email.SmtpServer, 25, email.Username, email.Passowrd)

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}