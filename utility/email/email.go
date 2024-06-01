package email

import (
	"gopkg.in/gomail.v2"
)

type SmtpConf struct {
	SmtpHost     string `json:"smtp_host"`
	SmtpPort     int    `json:"smtp_port"`
	SmtpUserName string `json:"smtp_username"`
	SmtpPassword string `json:"smtp_password"`

	Fromname string `json:"email_fromname"` //发送者的名字  默认发送者邮箱地址
}

// SendSMTPMail @SendSMTPMail   oss文件以附件形式发送到用户邮箱
// @tomMail   接收者邮箱地址  例如：15556566363@163.com
// @subject   邮件标题
// @body      邮件正文
// @attachName   邮件附件名字 例如：https://src/27840b463b3f36a2fbfcca5b5078420e.pdf
// @attachUrl    邮件附件oss地址
func SendSMTPMail(smtpConf *SmtpConf, toEmail, subject, body string) error {
	/*
		response, err := http.Get(attachUrl)
		if err != nil {
			return err
		}

		attach, err := io.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			return err
		}

	*/

	/*
		smtpConf := SmtpConf{
			SmtpHost:     "smtp.***.***.com", //邮箱服务地址  注意qq和163邮箱的也不一样（具体可以搜索：go 发送邮件 邮箱服务地址配置）
			SmtpPort:     465,                //端口  可以写成465
			SmtpUserName: "12******.cn",      //发送人的邮箱账号
			SmtpPassword: "******zu",         //发送人的邮箱密码
		}
	*/

	nickname := smtpConf.Fromname
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(smtpConf.SmtpUserName, nickname)) //发送者的名字  默认发送者邮箱地址
	m.SetHeader("To", toEmail)                                            //设置接收者邮箱
	m.SetHeader("Subject", subject)                                       //设置邮件标题
	if body != "" {
		m.SetBody("text/html", body) //设置邮件正文
	}

	/*
		m.Attach(attachName, gomail.SetCopyFunc(func(w io.Writer) error { //设置邮件附件  第一个参数是附件名称   第二个参数是二进制文件流
			_, err := w.Write(attach)
			return err
		}))

	*/
	d := gomail.NewDialer(smtpConf.SmtpHost, smtpConf.SmtpPort, smtpConf.SmtpUserName, smtpConf.SmtpPassword) //初始化配置
	return d.DialAndSend(m)                                                                                   //发送邮件
}
