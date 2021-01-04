package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"videoWeb/common/ecode"
	"videoWeb/notice-service/config"
)

var Not Service

type service struct {
	smtpStrategy map[string]smtpStrategy
}

type smtpStrategy struct {
	sender   string
	password string
	addr     string
	port     int
}

func New(conf *config.Config, logger log.Logger) Service {
	service := service{}
	service.smtpStrategy = make(map[string]smtpStrategy)
	service.smtpStrategy["QQ"] = smtpStrategy{
		sender:   conf.Email.QQ.Sender,
		password: conf.Email.QQ.Password,
		addr:     conf.Email.QQ.Addr,
		port:     conf.Email.QQ.Port,
	}
	return &service
}

type Service interface {
	SendEmail(ctx context.Context, strategyName string, title string, body string, recipientList []string) (ecode.ECode, error)
}
