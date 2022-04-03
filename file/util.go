package file

import (
	"fmt"
	"os"
	"io/ioutil"
)

func TextOpen(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("File Open Error!")
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)

	text := string(b) 
	return text
}
