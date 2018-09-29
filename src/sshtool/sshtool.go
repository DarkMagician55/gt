package sshtool

import (
	"github.com/DarkMagician55/terminal/src/common"
	"fmt"
	"encoding/hex"
)

const SEP = ":"

func CreatePwdStr(user, pwd string) (string, bool) {
	key := common.GetAesKey()
	input := user + SEP + pwd
	b , err := common.AesEncrypt([]byte(input), []byte(key))
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	return hex.EncodeToString(b), true
	//for test//for test//for test
	//for test
	//for test
	//for test
	//for test
	//for test
	//for test
}
