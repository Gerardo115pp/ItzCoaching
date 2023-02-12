package helpers

import (
	_ "image/png" // import PNG format for image.Decode
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

func PathExists(path_name string) bool {
	_, err := os.Stat(path_name)

	return !os.IsNotExist(err)
}

func WriteFile(file_name string, content []byte) error {
	file, err := os.Create(file_name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)

	return err
}
