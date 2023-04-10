package handler

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		/*
			上传文件到腾讯云
		*/
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		// 判断文件是否存在
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			fmt.Println("file.Read(b)")
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			fmt.Println("svcCtx.Engine.Where(\"hash = ?\", hash).Get(&rp)")
			return
		}
		if has {
			fmt.Println("文件已存在")
			httpx.OkJson(w, &types.FileUploadReply{
				Identity: rp.Identity,
				Ext:      rp.Ext,
				Name:     rp.Name,
			})
			return
		}
		// 往COS中存储数据
		cosPath, err := helper.CosUpload(r)
		if err != nil {
			fmt.Println("helper.CosUpload(r)")
			return
		}
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = cosPath
		/*
			----------over--------------
		*/
		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
