package homestay

import (
	"context"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBusinessAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessAddLogic {
	return &HomestayBusinessAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessAddLogic) HomestayBusinessAdd(req *types.AddHomestayBusinessReq) (resp *types.AddHomestayBusinessResp, err error) {
	// todo: add your logic here and delete this line

	return
}
