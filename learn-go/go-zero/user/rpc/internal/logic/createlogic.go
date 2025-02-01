package logic

import (
	"context"
	"database/sql"
	"time"

	"user/rpc/internal/svc"
	"user/rpc/models"
	"user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.CreateReq) (*user.CreateResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &models.User{
		Id: in.Id,
		// Avatar: in.Avatar,
		Name:  in.Name,
		Phone: in.Phone,
		CreatedAt: sql.NullTime{
			Time: time.Time{},
		},
		UpdatedAt: sql.NullTime{
			Time: time.Time{},
		},
	})
	return &user.CreateResp{}, err
}
