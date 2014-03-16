// Package ftp implements a FTP client as described in RFC 959.
package main

import (
	"fmt"
	goopt "github.com/droundy/goopt"
	ftp "github.com/jlaffaye/ftp"
)

var server_ip = goopt.String([]string{"-h"}, "0.0.0.0", "FTP服务器IP地址")
var server_port = goopt.Int([]string{"-p"}, 21, "FTP服务器端口")
var username = goopt.String([]string{"-u"}, "anonymous", "登陆用户名")
var password = goopt.String([]string{"-k"}, "anonymous", "登陆用户密码")

var dir = goopt.String([]string{"-d"}, "null", "所要查询的目录")
var file = goopt.String([]string{"-f"}, "null", "所要查询的文件名")

func main() {

	goopt.Description = func() string {
		return "Example program for using the goopt flag library."
	}
	goopt.Version = "1.0"
	goopt.Summary = "checker.exe -h 127.0.0.1 -u user -p 123qwe -d /dir -f file1"
	goopt.Parse(nil)

	fmt.Printf("FTP agrs info server_ip[%v]]\n", *server_ip)
	fmt.Printf("FTP agrs info server_port[%v]\n", *server_port)
	fmt.Printf("FTP agrs info username[%v]\n", *username)
	fmt.Printf("FTP agrs info password[%v]\n", *password)
	fmt.Printf("FTP agrs info dir[%v]\n", *dir)
	fmt.Printf("FTP agrs info file[%s]\n", *file)

	c, err := ftp.Connect("localhost:21")
	if err != nil {
		fmt.Printf("FTP Connect return : err\n")
		return
	}

	fmt.Printf("FTP Connect to server\n")

	err = c.Login(*username, *password)
	if err != nil {
		fmt.Printf("FTP Login return : ", err)
		return
	}

	fmt.Printf("FTP Login PASS!\n")

	err = c.ChangeDir(*dir)
	if err != nil {
		fmt.Printf("FTP ChangeDir return : ", err)
		return
	}

	fmt.Printf("FTP ChangeDir PASS!\n")

	entries := []*ftp.Entry{}

	entries, err = c.List(".")
	if err != nil {
		fmt.Printf("FTP List return : ", err)
		return
	}

	fmt.Printf("FTP List PASS!\n")

	for i, _ := range entries {

		fmt.Printf("    List Name[%s]\n", entries[i].Name)
		fmt.Printf("    List Type[%v]\n", entries[i].Type)
		fmt.Printf("    List Size[%d]\n", entries[i].Size)
		fmt.Printf("    List Time[%v]\n", entries[i].Time)

		if *file == entries[i].Name {
			fmt.Printf("OK , find the file[%s] in the FTP Server!\n", *file)

			c.Quit()
			fmt.Printf("FTP Quit!\n")
			return
		}
	}

	fmt.Printf("Sorry , find the file[%s] in the FTP Server!\n", *file)
	c.Quit()
	fmt.Printf("FTP Quit!\n")
	return
}
