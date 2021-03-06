package moni

import (
 "time"
)

type Host struct {
	Addr string
	Username string
	Password string
	Path string
	Commands []string
}

type hostitem struct {
	addr string
	username string
	password string
	commands []string
	path string
	sshcli   *SSHCli
	lastcheck time.Time
}

func (host*hostitem) GetCommands()[]string {
	return host.commands
}

func (host *hostitem) addCommand(command string){
	host.commands = append(host.commands, command)
}