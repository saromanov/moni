package moni

import(
"fmt"
)

//Useful commands for checking state of the system

const (
	Diskspace = "df"
)

type Outputfunc func(string)string

//DiskSpace provides parse result from df -h
func diskSpace(info string) string{
	fmt.Println("DISK SPACE")
	return ""
}