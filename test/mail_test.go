package test

import (
	"cloud-disk/core/define"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <getcharzhaopan@163.com>"
	//e.To = []string{toUserEmail}
	e.To = []string{"getcharzp@qq.com"}
	e.Subject = "验证码已发送，请查收"
	//e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	e.HTML = []byte("您的验证码：<b>123456</b>")
	//e.SendWithTLS("smtp.163.com:465",
	//	smtp.PlainAuth("", "getcharzhaopan@163.com", define.MailPassword, "smtp.163.com"),
	//	&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "949244762@qq.com", define.MailPassword, "smtp.qq.com"))
	if err != nil {
		t.Fatal(err)
	}
}
