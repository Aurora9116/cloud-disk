package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/models"
	"context"
	"time"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	uf := make([]*types.UserFile, 0)
	var cnt int64
	resp = new(types.UserFileListReply)
	size := req.Size
	if size == 0 {
		size = define.DefaultPageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	page = (page - 1) * size
	err = l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? and user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity,user_repository.repository_identity, user_repository.ext,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? or user_repository.deleted_at is null", time.Time{}.Format(define.Datetime)).
		Limit(size, page).Find(uf)
	if err != nil {
		return
	}
	cnt, err = l.svcCtx.Engine.Where("parent_id = ? and user_identity = ?", req.Id, userIdentity).Count(new(models.UserRepository))
	if err != nil {
		return
	}
	resp.List = uf
	resp.Count = cnt
	return
}
