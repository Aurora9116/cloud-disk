package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {
	// 判断当前名称是否存在
	count, err := l.svcCtx.Engine.Where("name = ? and parent_id ", req.Name, req.ParentId).Count(new(models.UserRepository))
	if err != nil {
		return
	}
	if count > 0 {
		return nil, errors.New("名称已存在")
	}
	// 创建文件夹
	data := &models.UserRepository{
		ParentId:     req.ParentId,
		Identity:     helper.GetUUID(),
		UserIdentity: userIdentity,
		Name:         req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return
	}
	return
}
