package homestay

import (
	"context"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayAddLogic {
	return &HomestayAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayAddLogic) HomestayAdd(req *types.AddHomestayReq) (resp *types.AddHomestayResp, err error) {
	// todo: add your logic here and delete this line

	return
}
