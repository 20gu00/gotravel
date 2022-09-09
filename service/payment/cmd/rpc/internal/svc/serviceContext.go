package svc

import (
	"go-travel/service/payment/cmd/rpc/internal/config"
	"go-travel/service/payment/model"

	"github.com/zeromicro/go-queue/kq"
)

type ServiceContext struct {
	Config                             config.Config
	ThirdPaymentModel                  model.ThirdPaymentModel
	KqueuePaymentUpdatePayStatusClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
