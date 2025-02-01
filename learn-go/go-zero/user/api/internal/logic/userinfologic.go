package logic

import (
	"context"

	"user/api/internal/svc"
	"user/api/internal/types"
	"user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	// todo: add your logic here and delete this line
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
