package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

func Md5(value string) string {
	hash := md5.New()
	return fmt.Sprintf("%x", hash.Sum([]byte(value)))
}
func GenerateToken(id int, identity string, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := claims.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", errors.New("claims SignedString error:" + err.Error())
	}
	return tokenString, nil
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
