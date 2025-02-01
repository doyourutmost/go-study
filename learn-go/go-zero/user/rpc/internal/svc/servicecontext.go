package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/rpc/internal/config"
	"user/rpc/models"
)

type ServiceContext struct {
	Config config.Config

	UserModel models.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,

		UserModel: models.NewUserModel(sqlConn, c.Cache),
	}
}
