package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"videoWeb/common/ecode"
	"videoWeb/customer-service/config"
	"videoWeb/customer-service/dao"
	genpb "videoWeb/generate-service/proto"
)

var Cus Service

type service struct {
	dao *dao.Dao
	gen genpb.GenerateServer
}

func New(conf *config.Config, dao *dao.Dao, gen genpb.GenerateServer, logger log.Logger) Service {
	return &service{dao: dao, gen: gen}
}

type Service interface {
	RegisterByPhone(ctx context.Context, phone string) (ecode.ECode, error)
	LoginByPhone(ctx context.Context, phone string) (ecode.ECode, string, error)
	GetCustomerInfoByToken(ctx context.Context, token string) (ecode.ECode, *dao.Customer, error)
}
