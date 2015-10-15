package moni

import (
	"fmt"
	"time"
	"errors"
	"log"
	"github.com/hashicorp/serf/serf"
)

type Moni struct {
	hosts    []string
	commands []string

	//List of remote hosts with usetname and passwords
	sshcli   []*SSHCli
	config   *Config
	serf     *serf.Serf
}

//New provides initialization of Moni
func New(path string) *Moni {
	moni := new(Moni)
	moni.config = LoadConfigData(path)
	moni.serf = serfInit()
	return moni
}

//AddNodes provides append nodes for monitoring
func (m *Moni) AddNodes(addr []string)(int, error){
	log.Printf("Add number of nodes: %d", len(addr))
	m.hosts = addr
	return m.serf.Join(addr, true)
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
	m.sshcli = initClients(m.config.Hosts)
	fmt.Printf("Start monitoring %s", time.Now().String())
	for {
		go func(commands []string) {
			for _, command := range commands {
				m.execute("default", command)
			}
		}(m.commands)

		time.Sleep(m.config.Timeout)
	}
}

//Execute current command
func (m *Moni) execute(host, command string) {
	for _, sshex := range m.sshcli {
		sshex.Exec(host, command)
	}
}

func (m *Moni) checkHosts()error {
	fmt.Println("CONFIG: ", m.config)
	if len(m.config.Hosts) == 0 {
		return errors.New("Information about hosts is not found")
	}

	return nil
}

func initClients(hosts []*Host)[]*SSHCli {
	result := []*SSHCli{}
	for _, host := range hosts {
		sshcli := NewSSHClient()
		sshcli.AuthUsernamePassword(host.Username, host.Password)
		result = append(result, sshcli)
	}
	return result
}

func serfInit()*serf.Serf {
	conf := serf.DefaultConfig()
	serfin, err := serf.Create(conf)
	if err != nil {
		log.Fatal(err)
	}

	return serfin
}
