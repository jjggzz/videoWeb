package service

import (
	"context"
	"gopkg.in/gomail.v2"
	"videoWeb/common/ecode"
	"videoWeb/common/ecode/business"
)

func (srv *service) SendEmail(ctx context.Context, strategyName string, title string, body string, recipientList []string) (ecode.ECode, error) {
	strategy, ok := srv.smtpStrategy[strategyName]
	if !ok {
		return business.SMTPStrategy, nil
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
