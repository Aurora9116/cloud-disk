package test

import (
	"cloud-disk/core/define"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"testing"
)

func TestFileUploadByFilepath(t *testing.T) {
	u, _ := url.Parse("https://aurora-1307772891.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})

	key := "cloud-disk/old.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./img/Screenshot_20210321_231810.jpg", nil,
	)
	if err != nil {
		panic(err)
	}
}
