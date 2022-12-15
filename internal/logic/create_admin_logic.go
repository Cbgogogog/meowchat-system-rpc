package logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"meowchat-notice-rpc/internal/model"
	"time"

	"meowchat-notice-rpc/internal/svc"
	"meowchat-notice-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAdminLogic {
	return &CreateAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAdminLogic) CreateAdmin(in *pb.CreateAdminReq) (*pb.CreateAdminResp, error) {
	id := primitive.NewObjectID()

	err := l.svcCtx.AdminModel.Insert(l.ctx, &model.Admin{
		ID:          id,
		CommunityId: in.CommunityId,
		Name:        in.Name,
		Title:       in.Title,
		Phone:       in.Phone,
		Wechat:      in.Wechat,
		AvatarUrl:   in.AvatarUrl,
		UpdateAt:    time.Time{},
		CreateAt:    time.Time{},
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateAdminResp{
		Id: id.Hex(),
	}, nil
}