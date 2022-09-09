package homestay

import (
	"context"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBusinessDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessDeleteLogic {
	return &HomestayBusinessDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessDeleteLogic) HomestayBusinessDelete(req *types.DeleteHomestayBusinessReq) (resp *types.DeleteHomestayBusinessResp, err error) {
	// todo: add your logic here and delete this line

	return
}
