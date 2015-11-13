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
	checks   []*Check
	hostlist []*hostitem
	//commands contans functions for formatting output
	outputfuncs map[string]*Command
}

//New provides initialization of Moni
func New(path string) *Moni {
	moni := new(Moni)
	moni.hosts = map[string]*hostitem{}
	moni.outputfuncs = map[string]*Command{}
	moni.config = LoadConfigData(path)
	moni.serf = serfInit()
	return moni
}

// AddNodes provides append nodes for monitoring.
// Nodes represent as Host object
func (m *Moni) AddNodes(hosts []*Host)(int, error){
	log.Printf("Add number of nodes: %d", len(hosts))
	if len(hosts) == 0 {
		return 0, fmt.Errorf("Node list is empty")
	}
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
//Event is a script
func (m *Moni) AddEvent(command string, item func(data string) bool) {
	
}

//AddMonitoring provides getting list of commands for monitoring
//list of commands available in commands.go
//Example AddMonitoring([]string{moni.Diskspace})
//Also, it can be loaded from config file
func (m *Moni) AddMonitoring(listcommands []string) {
	for _, command := range listcommands {
		switch command {
		case Diskspace:
			m.commands = append(m.commands, DiskSpaceCommand)
			m.outputfuncs[DiskSpaceCommand] = &Command {
				Title: "Free disk space",
				F: diskSpace,
			}

		case Networkinterfaces:
			m.commands = append(m.commands, NetworkinterfacesCommand)
			m.outputfuncs[NetworkinterfacesCommand] = &Command {
				Title: "Network interfaces",
				F: networkInterfaces,
			}
		}
	}
}

//AddChecks provides setting list of checks to monitoring
func (m *Moni) AddChecks(checks []*Check) error {
	if len(checks) == 0 {
		return fmt.Errorf("List of checks os empty")
	}
	m.checks = checks
	return nil
}

// Start provides starting of monitoring
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
				for _, command := range host.commands {
					result, err := m.execute(host.sshcli, host.addr, command)
					if err != nil {
						log.Print(err)
					}
					command, ok := m.outputfuncs[command]
					if ok {
						res, err := command.F(result)
						if err == nil {
							Show(fmt.Sprintf("%s:\n %s", command.Title, res))
							if m.config.Outpath != "" {
								Write(m.config.Outpath, fmt.Sprintf("%s:\n %s", command.Title, res))
							}
						}
					} else {
						Show(result)
					}
				}
			}
		}(m.hostlist)

		time.Sleep(m.config.Timeout)
	}
}

//Execute current command
func (m *Moni) execute(sshcli *SSHCli, host, command string) (string, error) {
	output, err := sshcli.Exec(host, command)
	if err != nil {
		return "", err
	}
	
	return output, nil
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
