package moni

import(
  "fmt"
  "time"
)

type Output struct {
	Data string
}

func (output*Output) Fit(){
	fmt.Printf("%s: %s", time.Now().String(), output.Data)
}