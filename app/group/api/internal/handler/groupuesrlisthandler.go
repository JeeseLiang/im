package handler

import (
	"net/http"

	"im_message/common/xresp"

	"im_message/app/group/api/internal/logic"
	"im_message/app/group/api/internal/svc"
	"im_message/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GroupUesrListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupUserListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := xresp.Validate.StructCtx(r.Context(), req); err != nil {
			xresp.Response(r, w, nil, err)
			return
		}

		l := logic.NewGroupUesrListLogic(r.Context(), svcCtx)
		resp, err := l.GroupUesrList(&req)
		xresp.Response(r, w, resp, err)
	}
}
