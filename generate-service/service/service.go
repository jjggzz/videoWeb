package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"videoWeb/generate-service/config"
)

var Gen Service

type service struct {
	node *snowflake.Node
}

func New(conf *config.Config) Service {
	node, err := snowflake.NewNode(conf.WorkId)
	if err != nil {
		panic(err)
	}
	return &service{node: node}
}

type Service interface {
	GenerateInt64Key(ctx context.Context) (int64, error)
	GenerateStringKey(ctx context.Context) (string, error)
}
