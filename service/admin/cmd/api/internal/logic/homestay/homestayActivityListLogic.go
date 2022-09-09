package homestay

import (
	"context"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayActivityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayActivityListLogic {
	return &HomestayActivityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayActivityListLogic) HomestayActivityList(req *types.ListHomestayActivityReq) (resp *types.ListHomestayActivityResp, err error) {
	// todo: add your logic here and delete this line

	return
}
