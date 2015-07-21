package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	filename := os.Args[1]
	data, _ := ioutil.ReadFile(filename)
	md5hash := md5.New()
	md5hash.Write(data)
	fmt.Println("md5: " + hex.EncodeToString(md5hash.Sum(nil)))
}
