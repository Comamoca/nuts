package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func RemoveFiles(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
