package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no filename provided")
		return
	}

	filename := os.Args[1]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("unable to read file")
		return
	}

	md5hash := md5.New()
	md5hash.Write(data)

	sha256hash := sha256.New()
	sha256hash.Write(data)

	fmt.Println("md5: " + hex.EncodeToString(md5hash.Sum(nil)))
	fmt.Println("sha256: " + hex.EncodeToString(sha256hash.Sum(nil)))
}
