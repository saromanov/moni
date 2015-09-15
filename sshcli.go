package main

import
(
	"golang.org/x/crypto/ssh"
	"log"
	"bytes"
)

type SSHCli struct {
	config *ssh.ClientConfig
}

func NewSSHClient(user, pass string)* SSHCli {
	sshcli := new(SSHCli)
	sshcli.config = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}

	return sshcli
}

//Exec provides execute command on the target host
func (sshcli*SSHCli) Exec(host, command string) string {
	conn, err := ssh.Dial("tcp", host+":22", sshcli.config)
	if err != nil {
		log.Fatal(err)
	}
	session, err2 := conn.NewSession()
	if err2 != nil {
		log.Fatal(err)
	}
	defer session.Close()
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(command)
	return stdoutBuf.String()
}