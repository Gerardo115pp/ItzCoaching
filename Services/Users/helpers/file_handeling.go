package helpers

import (
	"io"
	"os"
)

func CopyFile(from_name string, to_name string) (err error) {
	from, err := os.Open(from_name)
	if err != nil {
		return
	}
	defer from.Close()

	to, err := os.OpenFile(to_name, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return
}
