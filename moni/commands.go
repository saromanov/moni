package moni

//Useful commands for checking state of the system

const (
	Dikspace = "df"
)

type Outputfunc func(string)string

//DiskSpace provides parse result from df -h
func diskSpace(info string) string{
	return ""
}