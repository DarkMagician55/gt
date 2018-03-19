package sshtool

import (
	"os"
	"fmt"
	"bufio"
	"os/user"
)

func loadPwdStr() string {
	curUser, err := user.Current()
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
	fi, err := os.Open(curUser.HomeDir+"/.gs")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	a, _, _:= br.ReadLine()

	str := string(a)
	return str
}