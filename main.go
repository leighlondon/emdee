package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	md5Flag    bool
	sha256Flag bool
)

func init() {
	flag.BoolVar(&md5Flag, "m", false, "Calculate the MD5 hash.")
	flag.BoolVar(&sha256Flag, "s", false, "Calculate the SHA256 hash.")
}

func main() {

	// Parse the flags.
	flag.Parse()

	// If no flags were set, run everything.
	if flag.NFlag() == 0 {
		md5Flag = true
		sha256Flag = true
	}

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
