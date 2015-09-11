package main

import
(
	"golang.org/x/crypto/ssh"
	"os"
	"log"
	"bytes"
)

type SSHCli struct {
	config *ssh.ClientConfig
}

func NewSSHClient()* SSHCli {
	sshcli := new(SSHCli)
	sshcli.config = &ssh.ClientConfig{
		User: os.Getenv("LOGNAME"),
		Auth: []ssh.ClientAuth(keyring())
	}

	return sshcli
}

//Exec provides execute command on the target host
func (sshcli*SSHCli) Exec(host, command string) string {
	conn, err := ssh.Dial("tcp", hostname+":22", sshcli.config)
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
	session.Run(cmd)
	return stdoutBuf.String()
}

func keyring() ssh.ClientAuth {
	signers := []ssh.Singer{}
	keys := []string{os.Getenv("HOME") + "/.ssh/id_rsa", os.Getenv("HOME") + "/.ssh/id_dsa"}
	for _, keyname := range keys {
		singer, err := makeSinger(keyname)
		if err == nil {
			singers = append(singers, singer)
		}
	}

	return ssh.ClientAuthKeyring(&SingerContainer{singers})
}