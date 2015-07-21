package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
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

	if flag.NArg() == 0 {
		fmt.Println("no filename provided")
		return
	}

	for _, filename := range flag.Args() {

		// Load the data from file.
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(filename + ": unable to read file")
			return
		}

		// Declare the hashes.
		md5hash := md5.New()
		sha256hash := sha256.New()

		if md5Flag {
			md5hash.Write(data)
		}
		if sha256Flag {
			sha256hash.Write(data)
		}

		if md5Flag {
			fmt.Println("md5:    " + hex.EncodeToString(md5hash.Sum(nil)))
		}
		if sha256Flag {
			fmt.Println("sha256: " + hex.EncodeToString(sha256hash.Sum(nil)))
		}
	}
}
