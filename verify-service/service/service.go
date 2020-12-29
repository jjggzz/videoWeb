package service

import (
	"context"
	"videoWeb/common/ecode"
	"videoWeb/verify-service/dao"
)

var Ver Service

type service struct {
	dao *dao.Dao
}

func New(dao *dao.Dao) Service {
	return &service{dao: dao}
}

type Service interface {
	SendPhoneVerify(ctx context.Context, target string) (ecode.ECode, error)
	SendEmailVerify(ctx context.Context, target string) (ecode.ECode, error)
	CheckVerify(ctx context.Context, target string, code string) (ecode.ECode, error)
}
