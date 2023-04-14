package helper

import (
	"cloud-disk/core/define"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

func Md5(value string) string {
	hash := md5.New()
	return fmt.Sprintf("%x", hash.Sum([]byte(value)))
}

func GenerateToken(id int, identity string, name string, second int) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(second) * time.Second).Unix(),
			Issuer:    "",
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := claims.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", errors.New("claims SignedString error:" + err.Error())
	}
	return tokenString, nil
}

func AnalyzeToken(tokenString string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(tokenString, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claim, ok := claims.Claims.(*define.UserClaim); ok && claims.Valid {
		return claim, nil
	}
	return nil, errors.New("token analyze error ")
}

func MailSendCode(toUserEmail string, code string) error {
	e := email.NewEmail()
	e.From = "Get <949244762@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	//e.SendWithTLS("smtp.163.com:465",
	//	smtp.PlainAuth("", "getcharzhaopan@163.com", define.MailPassword, "smtp.163.com"),
	//	&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "949244762@qq.com", define.MailPassword, "smtp.qq.com"))
	if err != nil {

		return err
	}

	return nil
}

func RandCode() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < define.CodeLength; i++ {
		s = s + strconv.Itoa(rand.Intn(10))
	}
	return s
}

func GetUUID() string {
	return uuid.NewV4().String()
}

func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})

	file, fileHeader, err := r.FormFile("file")
	key := define.Disk + GetUUID() + path.Ext(fileHeader.Filename)
	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		panic(err)
	}

	return define.CosBucket + "/" + key, nil
}

func CosGet(name string) (string, error) {
	//u, _ := url.Parse(define.CosBucket)
	//b := &cos.BaseURL{BucketURL: u}
	//client := cos.NewClient(b, &http.Client{
	//	Transport: &cos.AuthorizationTransport{
	//		SecretID:  define.SecretID,
	//		SecretKey: define.SecretKey,
	//	},
	//})
	//key := define.Disk + name
	//objectURL := client.Object.GetObjectURL(key)
	//return objectURL.String(), nil
	// 替换成您的密钥
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})
	ak := define.SecretID
	sk := define.SecretKey
	name = strings.ReplaceAll(name, define.CosBucket+"/", "")
	ctx := context.Background()
	//_, err := client.Object.GetToFile(context.Background(), name, "./example.jpg", nil)
	//if err != nil {
	//	panic(err)
	//}
	presignedURL, err := client.Object.GetPresignedURL(ctx, http.MethodGet, name, ak, sk, time.Hour, nil)
	if err != nil {
		panic(err)
	}
	//resp, err := client.Object.Get(ctx, name, nil)
	//if err != nil {
	//	panic(err)
	//}
	//bs, _ := io.ReadAll(resp.Body)
	//defer resp.Body.Close()

	return presignedURL.String(), nil
}
