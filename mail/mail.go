package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

type SMTP struct {
	FromMail             string
	FromName             string
	Identity             string
	Username             string
	Password             string
	Host                 string
	Port                 int
	MaxMailContentLength int
	SSL                  bool
	StartTLS             bool
}

type Mail struct {
	Tos     []string
	Subject string
	Content string
}

func (m Mail) Validate(max int) error {
	if m.Tos == nil || len(m.Tos) == 0 {
		return fmt.Errorf("tos is blank")
	}

	if len(m.Subject) == 0 {
		return fmt.Errorf("subject is blank")
	}

	if len(m.Content) == 0 {
		return fmt.Errorf("content is blank")
	}

	if len(m.Content) > max {
		return fmt.Errorf("mail content too long: > %dB", max)
	}

	return nil
}

func NewSMTP(fromMail, fromName, username, password, host string, port int, ssl, startTLS bool) *SMTP {
	return &SMTP{
		FromMail:             fromMail,
		FromName:             fromName,
		Identity:             "",
		Username:             username,
		Password:             password,
		Host:                 host,
		Port:                 port,
		MaxMailContentLength: 102400,
		SSL:                  ssl,
		StartTLS:             startTLS,
	}
}

func (s *SMTP) Send(mail Mail) error {
	if err := mail.Validate(s.MaxMailContentLength); err != nil {
		return err
	}

	header := make(map[string]string)
	header["From"] = fmt.Sprintf("%s<%s>", s.FromName, s.FromMail)
	header["To"] = strings.Join(mail.Tos, ";")
	header["Subject"] = mail.Subject
	header["Content-Type"] = "text/html; charset=UTF-8"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message += "\r\n" + mail.Content

	auth := smtp.PlainAuth(
		s.Identity,
		s.Username,
		s.Password,
		s.Host,
	)

	err := send(
		fmt.Sprintf("%s:%d", s.Host, s.Port),
		auth,
		s.FromMail,
		mail.Tos,
		[]byte(message),
		s.SSL,
		s.StartTLS,
	)

	if err != nil {
		return fmt.Errorf("sendmail fail: %s", err)
	}

	return nil
}

func send(addr string, auth smtp.Auth, from string, tos []string, body []byte, ssl, startTLS bool) error {
	host := strings.Split(addr, ":")[0]

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	var c *smtp.Client
	var err error
	if ssl {
		conn, err := tls.Dial("tcp", addr, tlsconfig)
		if err != nil {
			return fmt.Errorf("tls dial %s fail: %s", addr, err)
		}

		c, err = smtp.NewClient(conn, host)
		if err != nil {
			return fmt.Errorf("new smtp client fail: %v", err)
		}
	} else {
		c, err = smtp.Dial(addr)
		if err != nil {
			return fmt.Errorf("dial %s fail: %v", addr, err)
		}

		if startTLS {
			err = c.StartTLS(tlsconfig)
			if err != nil {
				return fmt.Errorf("cannot starttls: addr: %s, err: %v", addr, err)
			}
		}
	}

	defer c.Close()

	if auth != nil {
		ok, _ := c.Extension("AUTH")

		if ok {
			err = c.Auth(auth)
			if err != nil {
				return fmt.Errorf("auth fail: %s", err.Error())
			}
		}
	}

	err = c.Mail(from)
	if err != nil {
		return err
	}

	for _, addr := range tos {
		if err = c.Rcpt(addr); err != nil {
			fmt.Printf("cannot add receiver: %s: %s", addr, err)
			continue
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(body)

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
