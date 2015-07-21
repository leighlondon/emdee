package main

import (
	"crypto/md5"
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

	fmt.Println("md5: " + hex.EncodeToString(md5hash.Sum(nil)))
}
