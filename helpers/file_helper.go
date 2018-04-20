package helpers

import (
	"io/ioutil"
)

func FileRead(path string) (string, error){
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
