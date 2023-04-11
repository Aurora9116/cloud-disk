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

var SecretID = "AKIDBey9rSr232tlX6Um2Jes9QZYJxVqzwfL"
var SecretKey = "c0DE1ZqaCe2PH1Xq3DLH9yhS7g291f5I"
var CosBucket = "https://aurora-1307772891.cos.ap-nanjing.myqcloud.com"
var Disk = "cloud-disk/"

// DefaultPageSize 分页默认参数
var DefaultPageSize = 20

// Datetime 时间格式
var Datetime = "2006-01-02 15:04:05b"
