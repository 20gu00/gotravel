package homestay

import (
	"context"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestaycommentAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestaycommentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestaycommentAddLogic {
	return &HomestaycommentAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestaycommentAddLogic) HomestaycommentAdd(req *types.AddHomestaycommentReq) (resp *types.AddHomestaycommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
