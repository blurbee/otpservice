package comms

import (
	"caluxor.com/api"
	"caluxor.com/util"
	"github.com/go-mail/mail"
)

type smtpSender struct {
	user       string
	password   string
	mailserver string
	email      string
	port       int
}

func SMTPInit(cfg *util.Config) (e api.EmailSender, err api.StatusCode) {
	s := new(smtpSender)
	s.password = cfg.GetSecret(cfg.EmailCfg.PasswordEnv)
	if s.password == "" {
		util.Error("No mail password credentials.")
		return nil, api.CONFIG_ERROR
	}
	s.mailserver = cfg.EmailCfg.MailServer
	s.email = cfg.EmailCfg.Email
	s.port = cfg.EmailCfg.Port
	s.user = cfg.EmailCfg.User

	return s, api.OK
}

func (s *smtpSender) SendEmail(dest string, subject string,
	message string) (err api.StatusCode) {
	m := mail.NewMessage()
	m.SetHeader("From", s.email)
	m.SetHeader("To", dest)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	// m.SetAddressHeader("Cc", "oliver.doe@example.com", "Oliver")
	// m.SetBody("text/html", "Hello <b>Kate</b> and <i>Noah</i>!")
	//m.Attach("lolcat.jpg")

	d := mail.NewDialer(s.mailserver, s.port, s.user, s.password)

	if er := d.DialAndSend(m); er != nil {
		util.Error("Mail send failed:", er)
		err = api.SEND_ERROR
		return
	}
	return api.OK
}
