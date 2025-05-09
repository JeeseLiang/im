package logic

import (
	"context"

	"im_message/app/msg/api/internal/svc"
	"im_message/app/msg/api/internal/types"
	"im_message/common/biz"
	"im_message/common/ctxdata"
	"im_message/proto/msg"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadRequest) (*types.UploadResponse, error) {
	var pbUploadRequest msg.UploadRequest
	err := copier.Copy(&pbUploadRequest, req)
	if err != nil {
		return nil, err
	}

	if pbUploadRequest.Uuid == "" {
		pbUploadRequest.Uuid = biz.GetUuid()
	}

	userId := ctxdata.GetUidFromCtx(l.ctx)
	pbUploadRequest.SenderId = userId
	pbUploadResponse, err := l.svcCtx.MsgRpc.Upload(l.ctx, &pbUploadRequest)
	if err != nil {
		return nil, err
	}
	var resp types.UploadResponse
	copier.Copy(&resp, pbUploadResponse)
	return &resp, nil
}
