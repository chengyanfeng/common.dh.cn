package utils

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"net/smtp"
)

func Mail(to string, subject string, body string) {
	if IsEmpty(to) || IsEmpty(subject) || IsEmpty(body) {
		Error("SendMail", to, subject, body)
		return
	}
	host := "smtp.exmail.qq.com"
	port := 465
	email := "support@datahunter.cn"
	password := "mRocker8"

	header := P{}
	header["From"] = "DataHunter" + "<" + email + ">"
	header["To"] = to
	subject = base64.StdEncoding.EncodeToString([]byte(subject))
	header["Subject"] = "=?UTF-8?B?" + subject + "?="
	header["Content-Type"] = "text/html; charset=UTF-8"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	auth := smtp.PlainAuth(
		"",
		email,
		password,
		host,
	)

	err := SendMailTls(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		email,
		[]string{to},
		[]byte(message),
	)

	if err != nil {
		Error(err)
	}
}

func SendMailTls(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {

	c, err := func(addr string) (*smtp.Client, error) {
		conn, err := tls.Dial("tcp", addr, nil)
		if err != nil {
			Error("SendMail", err)
			return nil, err
		}
		//分解主机端口字符串
		host, _, _ := net.SplitHostPort(addr)
		return smtp.NewClient(conn, host)
	}(addr)
	//create smtp client
	//c, err := dial(addr)
	if err != nil {
		Error("SendMail", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				Error("SendMail", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
