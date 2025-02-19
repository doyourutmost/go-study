// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package server

import (
	"context"

	"user/rpc/internal/logic"
	"user/rpc/internal/svc"
	"user/rpc/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// 定义rpc方法
func (s *UserServer) GetUser(ctx context.Context, in *user.GetUserReq) (*user.GetUserResp, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UserServer) Create(ctx context.Context, in *user.CreateReq) (*user.CreateResp, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *UserServer) Ping(ctx context.Context, in *user.Request) (*user.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
