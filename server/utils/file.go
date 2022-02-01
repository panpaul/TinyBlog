package utils

import (
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func WriteFile(filePath string, content []byte) error {
	return ioutil.WriteFile(filePath, content, 0644)
}

func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func FileTouch(filePath string) bool {
	_, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return false
	}
	return true
}
