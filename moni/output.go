package moni

//Output information

import(
  "fmt"
  "time"
  "os"
)

//Write output data to the file
func Write(path, data string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		return err
	
	}
	if _, err = f.WriteString(data); err != nil {
		return err
	}
	return nil

}

func Show(data string){
	fmt.Printf("%s:\n %s\n", time.Now().String(), data)
}