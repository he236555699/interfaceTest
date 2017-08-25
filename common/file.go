package common

import (
	"io/ioutil"
	"os"
)

func IsExist(path string) (flag bool) {
	if _, err := os.Stat(path); os.IsExist(err) {
		flag = true
	}

	return
}

func ReadAll(path string) (content string, err error) {
	if exist := IsExist(path); !exist {
		return
	}

	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	content = string(contentBytes[:])

	return
}
