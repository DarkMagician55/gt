package sshtool

import (

	"golang.org/x/crypto/ssh"
	"os"
	"golang.org/x/crypto/ssh/terminal"
)


type SshClient struct {
	ssh.Client
}

func (sc *SshClient)Open() {

	session, err := sc.NewSession()
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