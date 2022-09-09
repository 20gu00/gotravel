package homestay

import (
	"context"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBusinessUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessUpdateLogic {
	return &HomestayBusinessUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessUpdateLogic) HomestayBusinessUpdate(req *types.UpdateHomestayBusinessReq) (resp *types.UpdateHomestayBusinessResp, err error) {
	// todo: add your logic here and delete this line

	return
}
