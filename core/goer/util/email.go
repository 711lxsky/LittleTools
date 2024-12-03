package util

import (
	"goer/config"
	"gopkg.in/gomail.v2"
	"log"
	"net"
	"net/smtp"
	"regexp"
	"strings"
)

func CheckEmailValid(email string) bool {
	// 正则表达式校验
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return false
	}
	// 域名验证，使用net包进行DNS查询，检查是否存在有效的MX记录
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	var mx string
	if mxRecords, err := net.LookupMX(domain); err != nil || len(mxRecords) <= 0 {
		return false
	} else {
		mx = mxRecords[0].Host
	}
	// SMTP验证
	conn, err := net.Dial("tcp", mx+":25")
	if err != nil {
		log.Fatalf("failed to connect to SMTP server: %v", err)
		return false
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}(conn)
	client, err := smtp.NewClient(conn, mx)
	if err != nil {
		log.Fatalf("failed to create SMTP client: %v", err)
		return false
	}
	defer func(client *smtp.Client) {
		err := client.Quit()
		if err != nil {
			log.Fatalf("failed to quit SMTP client: %v", err)
		}
	}(client)
	// 发送HELO命令 告诉邮件服务器我们的主机名
	if err := client.Hello("localhost"); err != nil {
		log.Printf("failed to send HELO command: %v", err)
		return false
	}

	// 发送MAIL FROM命令 指定发件人地址
	if err := client.Mail("3034047525@qq.com"); err != nil {
		log.Printf("failed to send MAIL FROM command: %v", err)
		return false
	}

	// 发送RCPT TO命令 指定收件人地址
	if err := client.Rcpt(email); err != nil {
		log.Printf("failed to send RCPT TO command: %v", err)
		return false
	}

	return true
}

func BuildDialer() *gomail.Dialer {
	return gomail.NewDialer(config.EmailSendHost, config.EmailSendPort, config.EmailSendUser, config.EmailSendPass)
}
