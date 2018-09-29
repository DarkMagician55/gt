package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/DarkMagician55/terminal/src/sshtool"
)

var createPwd  = flag.Bool("c", false, "Use -c -u user -p pwd")
var u          = flag.String("u", "", "Use -c <user>")
var pwd        = flag.String("p", "", "Use -c <pwd>")

func main() {

	flag.Parse()
	if flag.NArg() < 1 && !*createPwd{
		return
	}

	if *createPwd{
		if *u != "" && *pwd != "" {
			pwdStr , ok := sshtool.CreatePwdStr(*u, *pwd)
			if ok {
				fmt.Println(pwdStr)
				return
			}
		}
		fmt.Println("Use -c -u user -p pwd")
		return
	}

	host := flag.Arg(0)
	su := sshtool.NewSshUser()
	sc := su.Connect(host)
	if sc == nil {
		fmt.Println("connect fail")
		os.Exit(1)
	}
	sc.Open()

	//for test
	//for test
	//for test
	//for test
	//for test
	//for test
	//for test
	//for test
	//for test
	//for test
	//for test
	//for test
	//for test
	//for test
	//for test
	//for test

}