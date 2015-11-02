package moni

import (
	"fmt"
	"strings"
)

//Useful commands for checking state of the system

const (
	Diskspace         = "df"
	Networkinterfaces = "ifconfig"
	DiskSpaceCommand = "df -h"
	NetworkinterfacesCommand = "ifconfig -s"
)

type Outputfunc func(string) (string, error)

type Command struct {
	F     Outputfunc
	Title string
}

//DiskSpace provides parse result from df -h
func diskSpace(info string) (string, error) {
	line := strings.Split(info, "\n")[1]
	if len(line) == 0 {
		return "", fmt.Errorf("Can't get information about disk space")
	}

	fields := strings.Fields(line)
	if len(fields) < 4 {
		return "", fmt.Errorf("Can't get information about disk space")
	}
	return fields[3], nil
}

//networkInterfaces provides parse result from ifconfig
func networkInterfaces(info string) (string, error) {
	result :=""
	lines := strings.Split(info, "\n")
	for _, line := range lines[1:] {
		item := strings.Fields(line)
		if len(line) > 0 {
			result += fmt.Sprintf("Name: %s, bytes: %s\n", item[0], item[3])
		}
	}
	return result, nil
}
