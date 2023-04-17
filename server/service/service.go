package service

import (
	"context"
	"wukong/pkg/dal"
)

// 接口逻辑实现在service 中实现

type Service struct {
	ctx context.Context
}

func NewService() *Service {
	return &Service{
		ctx: context.Background(),
	}
}

var dalf = func() dal.Client {
	return dal.NewDal()
}
