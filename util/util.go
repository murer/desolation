package util

import (
	"io"
	"io/ioutil"
	"os"
)

var Version = "dev"

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadMax(r io.Reader, max int) []byte {
	ret := make([]byte, max)
	n, err := r.Read(ret)
	Check(err)
	return ret[:n]
}

func ReadMaxString(r io.Reader, max int) string {
	return string(ReadMax(r, max))
}

func ReadAll(r io.Reader) []byte {
	ret, err := ioutil.ReadAll(r)
	Check(err)
	return ret
}

func ReadAllString(r io.Reader) string {
	return string(ReadAll(r))
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	Check(err)
	return true
}
