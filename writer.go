package file

import (
	"os"
	"path"
	"io/ioutil"
)

func WriteBytes(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

func WriteString(filePath string, s string) (int, error) {
	return WriteBytes(filePath, []byte(s))
}

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	if perm == 0 { perm = os.ModePerm }
	if err := ioutil.WriteFile(filePath, data, perm); err != nil {
		return err
	}

	return nil
}