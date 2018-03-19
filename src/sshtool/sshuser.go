package sshtool

import (
	"github.com/DarkMagician55/terminal/src/common"
	"fmt"
	"encoding/hex"
	"strings"
	"golang.org/x/crypto/ssh"
)


type SshUser struct {
	user string
	pwd string
}

func NewSshUser() (*SshUser) {
	pwdStr := loadPwdStr()
	b, err := hex.DecodeString(pwdStr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	b, err = common.AesDecrypt(b, []byte(common.GetAesKey()))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	strArr := strings.Split(string(b), SEP)
	user := strArr[0]
	pwd := strings.Join(strArr[1:], "")

	su := SshUser{user,pwd}
	return &su
}

func NewSshUserByPwdStr(pwdStr string) (*SshUser) {
	b, err := hex.DecodeString(pwdStr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	b, err = common.AesDecrypt(b, []byte(common.GetAesKey()))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	strArr := strings.Split(string(b), SEP)
	user := strArr[0]
	pwd := strings.Join(strArr[1:], "")

	su := SshUser{user,pwd}
	return &su
}

func (su *SshUser)Connect(host string) (*SshClient) {
	client, err := ssh.Dial("tcp", host+":22", &ssh.ClientConfig{
		User: su.user,
		Auth: []ssh.AuthMethod{ssh.Password(su.pwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &SshClient{*client}
}