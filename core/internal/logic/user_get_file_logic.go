package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGetFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserGetFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGetFileLogic {
	return &UserGetFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserGetFileLogic) UserGetFile(req *types.UserGetFileRequest) (resp *types.UserGetFileReply, err error) {
	// todo: add your logic here and delete this line
	rp := new(models.RepositoryPool)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.Identity).Get(rp)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("文件不存在")
	}
	url, err := helper.CosGet(rp.Path)
	if err != nil {
		return nil, err
	}
	resp = new(types.UserGetFileReply)
	resp.UrL = url
	return
}
