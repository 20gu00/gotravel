package homestay

import (
	"context"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestaycommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestaycommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestaycommentListLogic {
	return &HomestaycommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestaycommentListLogic) HomestaycommentList(req *types.ListHomestaycommentReq) (resp *types.ListHomestaycommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
