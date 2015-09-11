package main

import (
	"fmt"
	"time"
)

type Moni struct {
	hosts    []string
	commands []string
	sshcli   *SSHCli
	config   *Config
}

//New provides initialization of Moni
func New(path string) *Moni {
	moni := new(Moni)
	moni.sshcli = NewSSHClient()
	moni.hosts = []string{}
	moni.config = LoadConfigData(path)
	return moni
}

//AddCommand provides append command for monitoring
func (m *Moni) AddCommand(command string) {
	m.command = append(m.command, command)
}

//AddEvent provides append new event
func (m *Moni) AddEvent(command string, item func(data string) bool) {

}

//Sart provides starting of monitoring
func (m *Moni) Start() {
	fmt.Printf("Start monitoring %s", time.String())
	for {
		go func(commands []string) {
			for _, command := range commands {
				m.execute(command)
			}
		}(m.commands)

		time.Sleep(m.config.timeout)
	}
}

//Execute current command
func (m *Moni) execute(command string) {
	for _, host := range m.hosts {
		m.sshcli.Exec(host, command)
	}
}
