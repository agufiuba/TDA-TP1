package utils

import (
	"io/ioutil"
)

func SaveToFile(fn, s string) error {
	bytes := []byte(s)
	err := ioutil.WriteFile(fn, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
