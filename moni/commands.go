package moni

import (
	"fmt"
	"strings"
)

//Useful commands for checking state of the system

const (
	Diskspace = "df"
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
