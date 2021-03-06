package util

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
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

func ReadFully(r io.Reader, n int) []byte {
	buf := make([]byte, n)
	nr, err := io.ReadAtLeast(r, buf, n)
	Check(err)
	if nr != n {
		log.Panicf("wrong %d, expected %d", nr, n)
	}
	return buf
}

func Rand(len int) []byte {
	token := make([]byte, len)
	n, err := rand.Read(token)
	Check(err)
	if n != len {
		log.Fatalf("wrong, exp: %d, but was: %d", len, n)
	}
	return token
}
