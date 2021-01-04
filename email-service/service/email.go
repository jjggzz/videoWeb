package service

import (
	"context"
	"gopkg.in/gomail.v2"
	"videoWeb/common/ecode"
)

func (srv *service) SendEmail(ctx context.Context, smtp string, title string, body string, recipientList []string) (ecode.ECode, error) {
	strategy, ok := srv.smtpStrategy[smtp]
	if !ok {
		return ecode.SMTPStrategy, nil
	}
	m := gomail.NewMessage()
	m.SetHeader("From", strategy.sender)
	m.SetHeader("To", recipientList...)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", body)
	err := gomail.NewDialer(strategy.addr, strategy.port, strategy.sender, strategy.password).DialAndSend(m)
	if err != nil {
		return ecode.ServerErr, err
	}
	return ecode.Success, nil
}
