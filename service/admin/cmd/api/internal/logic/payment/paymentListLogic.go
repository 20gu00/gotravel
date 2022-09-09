package payment

import (
	"context"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentListLogic {
	return &PaymentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentListLogic) PaymentList(req *types.ListPaymentReq) (resp *types.ListPaymentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
