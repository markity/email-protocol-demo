package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 25端口: 建立明文tcp连接, 可能被一些ISP拦截, 不建议使用
// 587端口: 同25, 但是不会被ISP拦截 nc smtp.qq.com 587
// 		加密: 首先通过587建立明文连接, 再通过STARTTLS升级成SSL/TLS连接
// 465端口: 直接建立SSL/TLS连接, 有可能邮件服务器没有实现这个选项 openssl s_client -connect smtp.qq.com

// 综上, 处于兼容性和安全性考虑, 一般用587端口+STARTTLS发送邮件
//		还有一个好处是, 如果邮件服务器没有实现STARTTLS, 程序也能选择用非加密的方式发出

func main() {
	conn, err := net.Dial("tcp", "smtp.qq.com:587")
	if err != nil {
		panic(err)
	}

	// conn reader
	go func() {
		for {
			buf := make([]byte, 256)
			n, err := conn.Read(buf)
			if err != nil {
				panic(err)
			}

			_, err = os.Stdout.Write(buf[:n])
			if err != nil {
				panic(err)
			}
		}
	}()

	// stdin reader qrcchbkvabeqbadb
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, err := conn.Write([]byte(fmt.Sprintf("%s\r\n", scanner.Text())))
		if err != nil {
			panic(err)
		}
	}

	select {}
}
