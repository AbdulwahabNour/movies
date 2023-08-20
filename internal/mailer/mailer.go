package mailer

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	"github.com/AbdulwahabNour/movies/config"
	"github.com/go-mail/mail/v2"
)

var (
	//go:embed templates/*
	templateFs embed.FS
)

type Mailer struct {
	dialer          *mail.Dialer
	sender          string
	templateDirName string
}

type MailerData struct {
	Data      map[string]interface{}
	Recipient string
}

func NewMailer(config *config.Config, templateDirName string) Mailer {

	dialer := mail.NewDialer(config.Mail.Host, config.Mail.Port, config.Mail.UserName, config.Mail.Password)
	dialer.Timeout = config.Mail.TimeOut

	return Mailer{
		dialer:          dialer,
		sender:          config.Mail.Sender,
		templateDirName: templateDirName,
	}
}
func (m Mailer) Send(data MailerData) error {
	tmpl, err := m.loadMailTemplate()
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return err
	}

	msg := mail.NewMessage()
	msg.SetHeader("To", data.Recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	err = m.dialer.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}
func (m Mailer) loadMailTemplate() (*template.Template, error) {

	ts, err := template.ParseFS(templateFs, fmt.Sprintf("templates/%s/*.tmpl", m.templateDirName))
	if err != nil {
		return nil, err
	}
	return ts, nil

}
