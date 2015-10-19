package moni

import
(
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	//"log"
	"bytes"
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

//AuthWithFile provides auth with /.ssh/id_rsa file
func (sshcli*SSHCli) AuthWithFile(username, path string) {
	/*data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	pubkey, err := ssh.ParsePrivateKey(data)
	if err != nil {
		log.Fatal(err)
	}*/

	sshcli.config = &ssh.ClientConfig {
		User: username,
		Auth: []ssh.AuthMethod{PublicKeyFile(path)},
	}

}

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

//Exec provides execute command on the target host
//Return result from command
func (sshcli*SSHCli) Exec(host, command string) (string, error) {
	host = "127.0.0.1"
	conn, err := ssh.Dial("tcp", host+":2667", sshcli.config)
	if err != nil {
		return "", err
	}
	session, err2 := conn.NewSession()
	if err2 != nil {
		return "", err2
	}
	defer session.Close()
	//result := &SSHResult{}

	var b bytes.Buffer
	session.Stdout = &b

	//stdout, errstdout := session.StdoutPipe()
	/*if errstdout != nil {
		return nil, errstdout
	}
	go io.Copy(result.output, stdout)

	stderr, errpos := session.StderrPipe()
	if errpos != nil {
		return nil, errpos
	}
	go io.Copy(result.stderr, stderr)*/
	errrun := session.Run(command)
	if errrun != nil {
		return "", errrun
	}
	return b.String(), nil
}