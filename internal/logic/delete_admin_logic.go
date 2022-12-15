package logic

import (
	"context"
	"meowchat-notice-rpc/errorx"

	"meowchat-notice-rpc/internal/svc"
	"meowchat-notice-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAdminLogic {
	return &DeleteAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAdminLogic) DeleteAdmin(in *pb.DeleteAdminReq) (*pb.DeleteAdminResp, error) {
	err := l.svcCtx.AdminModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, errorx.Switch(err)
	}

	return &pb.DeleteAdminResp{}, nil
}
