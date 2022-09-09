package homestay

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-travel/service/admin/cmd/api/internal/logic/homestay"
	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"
)

func HomestayActivityListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListHomestayActivityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homestay.NewHomestayActivityListLogic(r.Context(), svcCtx)
		resp, err := l.HomestayActivityList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
