package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/emailplus/utils"
)

var Email = new(email)

type email struct{}

//@author: [maplepie](https://github.com/maplepie)
//@function: EmailTest
//@description: 发送邮件测试
//@return: err error

func (s *email) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}

//@author: [maplepie](https://github.com/maplepie)
//@function: EmailTest
//@description: 发送邮件测试
//@return: err error
//@params to string 	 收件人
//@params subject string   标题（主题）
//@params body  string 	 邮件内容

func (s *email) SendEmail(to, subject, body string) (err error) {
	err = utils.Email(to, subject, body)
	return err
}
