package moni

// Check represent commands, which returns 0 if check is complete
// And 1 in otherwise
type Check struct {
	Path string
	Completed int
	Failed int
	PercentCompleted float64
}

func (ch*Check) Complete() {
	ch.Completed++
}

func(ch*Check) Fail() {
	ch.Failed++
}