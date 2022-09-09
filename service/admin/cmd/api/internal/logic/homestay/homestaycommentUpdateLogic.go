package homestay

import (
	"context"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestaycommentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestaycommentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestaycommentUpdateLogic {
	return &HomestaycommentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestaycommentUpdateLogic) HomestaycommentUpdate(req *types.UpdateHomestaycommentReq) (resp *types.UpdateHomestaycommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
