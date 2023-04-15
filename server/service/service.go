package service

import "context"

// 接口逻辑实现在service 中实现

type Service struct {
	ctx context.Context
}

func NewService() *Service {
	return &Service{
		ctx: context.Background(),
	}
}
