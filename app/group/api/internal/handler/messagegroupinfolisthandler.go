package handler

import (
	"net/http"

	"im_message/common/xresp"

	"im_message/app/group/api/internal/logic"
	"im_message/app/group/api/internal/svc"
	"im_message/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MessageGroupInfoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageGroupInfoListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := xresp.Validate.StructCtx(r.Context(), req); err != nil {
			xresp.Response(r, w, nil, err)
			return
		}

		l := logic.NewMessageGroupInfoListLogic(r.Context(), svcCtx)
		resp, err := l.MessageGroupInfoList(&req)
		xresp.Response(r, w, resp, err)
	}
}
