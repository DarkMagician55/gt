package common

import (
	"net"
	"time"
	"math/rand"
	"fmt"
	"crypto/md5"
)

//每次编译都不同
const salt  = "474a731c039a8c4588b595f821887f33"

func GetAesKey() (key string) {
	key += getMacAddr()
	key += salt
	return getMd5Str([]byte(key))
}

func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, inter := range interfaces {
		addr = inter.HardwareAddr.String()
		if addr != ""{
			return
		}
	}
	return
}

//生成随机字符串
func GetRandomString(length int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	l := len(bytes)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(l)])
	}
	return string(result)
}

func getMd5Str(in []byte) (out string) {
	h := md5.New()
	h.Write(in)
	out = fmt.Sprintf("%x", h.Sum(nil))
	return
}

