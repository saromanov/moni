package main

import (
	"fmt"
	"time"
	"errors"
	"log"
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
	moni.config = LoadConfigData(path)
	return moni
}

//AddCommand provides append command for monitoring
func (m *Moni) AddCommand(command string) {
	m.commands = append(m.commands, command)
}

//AddEvent provides append new event
func (m *Moni) AddEvent(command string, item func(data string) bool) {

}

//Sart provides starting of monitoring
func (m *Moni) Start() {
	err := m.checkHosts()
	if err != nil {
		log.Fatal(err)
	}

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

func (m *Moni) checkHosts()error {
	if len(m.config.Hosts) == 0 {
		return errors.New("Information about hosts is not found")
	}

	return nil
}
