package main

import (
	"golang.org/x/crypto/ssh"
	"fmt"
	"os"
	"runtime/debug"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			fmt.Println("panic!", err)
			os.Exit(1)
		}
	}()

	ce := func(err error, msg string) {
		if err != nil {
			fmt.Println(msg, err)
			os.Exit(1)
		}
	}

	host := ""
	pwd := ""

	client, err := ssh.Dial("tcp", host, &ssh.ClientConfig{
		User: "guotong",
		Auth: []ssh.AuthMethod{ssh.Password(pwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	ce(err, "dial")

	session, err := client.NewSession()
	ce(err, "new session")
	defer session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, oldState)

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := terminal.GetSize(fd)

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		//ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("xterm-256color", termHeight, termWidth, modes)
	ce(err, "request pty")

	err = session.Shell()
	ce(err, "start shell")

	//go no(session)

	err = session.Wait()
	ce(err, "return")
}