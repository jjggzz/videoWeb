package service

import (
	"context"
	"github.com/go-kit/kit/log"
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
	RegisterByPhone(ctx context.Context, phone string) (RegisterStatus, error)
}

type RegisterStatus int32

const (
	RegisterStatus_SUCCESS      = 0
	RegisterStatus_FAIL_HAS_USE = 1
	RegisterStatus_FAIL_SYS_ERR = -1
)
