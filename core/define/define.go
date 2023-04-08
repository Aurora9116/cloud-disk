package define

import (
	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"
var MailPassword = "qtssdidxkuytbcah"

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间(s)
var CodeExpire = 300
