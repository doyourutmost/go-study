package logic

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"
	"user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	getUserResp, err := l.svcCtx.User.GetUser(l.ctx, &userclient.GetUserReq{
		Id: req.Id,
	}) // rpc call
	if err != nil {
		return nil, err
	}
	resp = &types.GetUserResponse{
		Id:    getUserResp.Id,
		Name:  getUserResp.Name,
		Phone: getUserResp.Phone,
	}
	return
}
