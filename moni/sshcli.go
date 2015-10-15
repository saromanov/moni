package moni

import
(
	"golang.org/x/crypto/ssh"
	"io"
)


type SSHCli struct {
	config *ssh.ClientConfig
}

type SSHResult struct{
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
func (sshcli*SSHCli) AuthUsernamePassword(username, password string) {
	sshcli.config = &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
	}
}

//Exec provides execute command on the target host
//Return result from command
func (sshcli*SSHCli) Exec(host, command string) (*SSHResult, error) {
	conn, err := ssh.Dial("tcp", host+":22", sshcli.config)
	if err != nil {
		return nil, err
	}
	session, err2 := conn.NewSession()
	if err2 != nil {
		return nil, err2
	}
	defer session.Close()
	result := &SSHResult{}
	stdout, errstdout := session.StdoutPipe()
	if errstdout != nil {
		return nil, errstdout
	}
	go io.Copy(result.output, stdout)

	stderr, errpos := session.StderrPipe()
	if errpos != nil {
		return nil, errpos
	}
	go io.Copy(result.stderr, stderr)
	errrun := session.Run(command)
	if errrun != nil {
		return nil, errrun
	}
	return result, nil
}