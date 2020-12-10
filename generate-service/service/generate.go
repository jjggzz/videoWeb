package service

import (
	"context"
)

// 生成int64的key
func (srv *service) GenerateInt64Key(ctx context.Context) (int64, error) {
	return srv.node.Generate().Int64(), nil
}

// 生成string的key
func (srv *service) GenerateStringKey(ctx context.Context) (string, error) {
	return srv.node.Generate().String(), nil
}
