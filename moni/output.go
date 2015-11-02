package moni

//Output information

import(
  "fmt"
  "time"
)

func Show(data string){
	fmt.Printf("%s:\n %s\n", time.Now().String(), data)
}