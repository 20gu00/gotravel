package order

import (
	"context"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderAddLogic {
	return &OrderAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderAddLogic) OrderAdd(req *types.AddOrderReq) (resp *types.AddOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
