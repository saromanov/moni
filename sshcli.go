package main

import
(
	"golang.org/x/crypto/ssh"
	"log"
	"bytes"
	"io"
)


type SSHCli struct {
	config *ssh.ClientConfig
}

type SSHResult {
	command string
	output io.Writer
	stderr io.Writer
}

func NewSSHClient()* SSHCli {
	sshcli := new(SSHCli)
	sshcli.config = &ssh.ClientConfig{}
	return sshcli
}

//AuthUsernamePassword provides auth with password to ssh server 
func (sshcli*SSHCli) AuthUsernamePassword(username, passord string) {
	sshcli.config = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
}

//Exec provides execute command on the target host
//Return result from command
func (sshcli*SSHCli) Exec(host, command string) *SSHResult{
	conn, err := ssh.Dial("tcp", host+":22", sshcli.config)
	if err != nil {
		return nil, err
	}
	session, err2 := conn.NewSession()
	if err2 != nil {
		return nil, err
	}
	defer session.Close()
	result := &SSHResult{}
	stdout, errstdout := session.StdoutPipe()
	if errstdout != nil {
		return nil, err
	}
	go io.Copy(result.output, stdout)

	stderr, errpos := session.StderrPipe()
	if errpos != nil {
		return nil, err
	}
	go io.Copy(result.stderr, stderr)
	err := session.Run(command)
	if err != nil {
		return nil, err
	}
	return result, nil
}