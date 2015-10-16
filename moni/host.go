package moni

type Host struct {
	Addr string
	Username string
	Password string
	Commands []string
}

func (host*Host) GetCommands()[]string {
	return host.Commands
}

func (host *Host) AddCommand(command string){
	host.Commands = append(host.Commands, command)
}