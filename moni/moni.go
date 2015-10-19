package moni

import (
	"fmt"
	"time"
	"errors"
	"log"
	"net"
	"github.com/hashicorp/serf/serf"
)

type Moni struct {
	commands []string

	//List of remote hosts with usetname and passwords
	sshcli   []*SSHCli
	config   *Config
	serf     *serf.Serf
	hosts    map[string]*hostitem
	hostlist []*hostitem
}

//New provides initialization of Moni
func New(path string) *Moni {
	moni := new(Moni)
	moni.hosts = map[string]*hostitem{}
	moni.config = LoadConfigData(path)
	moni.serf = serfInit()
	return moni
}

//AddNodes provides append nodes for monitoring
func (m *Moni) AddNodes(hosts []*Host)(int, error){
	log.Printf("Add number of nodes: %d", len(hosts))
	m.config.Hosts = hosts
	addrs := make([]string, len(hosts))
	for i, host := range hosts {
		addrs[i] = host.Addr
		client, err := initClient(host)
		if err != nil {
			continue
		}
		item := &hostitem{
			addr: host.Addr,
			username: host.Username,
			password: host.Password,
			commands: host.Commands,
			sshcli: client,
			lastcheck: time.Now(),
		}

		m.hosts[host.Addr] = item
		m.hostlist = append(m.hostlist, item)
		//m.hosts[host.Addr] = append(m.hosts[host.Addr], item)
	}
	return m.serf.Join(addrs, true)
}

//AddCommand provides append command for monitoring for all hosts
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
	m.mergeCommands()
	fmt.Printf("Start monitoring %s", time.Now().String())
	for {
		go func(hosts []*hostitem) {
			for _, host := range hosts {
				m.execute(host.sshcli, host.addr, host.commands[0])
			}
		}(m.hostlist)

		time.Sleep(m.config.Timeout)
	}
}

//Execute current command
func (m *Moni) execute(sshcli *SSHCli, host, command string) {
	fmt.Println(sshcli.Exec(host, command))
}


func (m *Moni) checkHosts()error {
	if len(m.config.Hosts) == 0{
		return errors.New("Information about hosts is not found")
	}

	return nil
}

//checking availability of hosts
func (m *Moni) checkHostAvailibility(addr string) error {
	conn, err := net.DialTimeout("tcp", addr, 15 * time.Second)
	defer conn.Close()
	if err != nil {
		return err
	}

	return nil
}

//move commands from "commands" to all host. Because these commands
//should be executed on all of hosts
func (m *Moni) mergeCommands() {
	if len(m.commands) == 0 {
		return
	}

	for _, host := range m.hosts {
		for _, command := range m.commands {
			host.addCommand(command)
		}
	}
}

func initClient(host *Host)(*SSHCli, error) {
	sshcli := NewSSHClient()
	if host.Password != "" {
		sshcli.AuthUsernamePassword(host.Username, host.Password)
		return sshcli, nil
	} else if host.Path != "" {
		sshcli.AuthWithFile(host.Username, host.Path)
		return sshcli, nil
	}
	return sshcli, fmt.Errorf("SSH client is not created")
}

func serfInit()*serf.Serf {
	conf := serf.DefaultConfig()
	serfin, err := serf.Create(conf)
	if err != nil {
		log.Fatal(err)
	}

	return serfin
}
