package logic

import (
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	// 判断当前名称是否存在
	count, err := l.svcCtx.Engine.Where("name = ? and parent_id = (select parent_id form user_repository ur where ur.identity = ?)", req.Name, req.Identity).Count(new(models.UserRepository))
	if err != nil {
		return
	}
	if count > 0 {
		return nil, errors.New("名称已存在")
	}
	data := &models.UserRepository{
		Name: req.Name,
	}
	_, err = l.svcCtx.Engine.Where("identity = ? and user_identity = ?", req.Identity, userIdentity).Update(data)
	if err != nil {
		return
	}
	return
}
