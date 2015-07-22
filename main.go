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
	md5Flag     bool
	sha256Flag  bool
	versionFlag bool
)

func init() {
	flag.BoolVar(&md5Flag, "m", false, "Calculate the MD5 hash.")
	flag.BoolVar(&sha256Flag, "s", false, "Calculate the SHA256 hash.")
	flag.BoolVar(&versionFlag, "v", false, "Show the version number.")
}

func main() {

	// Parse the flags.
	flag.Parse()

	// Check for the easy flag.
	if versionFlag {
		fmt.Println(VersionString)
		return
	}

	// If no flags were set, run everything.
	if flag.NFlag() == 0 {
		md5Flag = true
		sha256Flag = true
	}

	// Need at least one filename to be provided.
	if flag.NArg() == 0 {
		fmt.Println("no filename provided")
		return
	}

	for _, filename := range flag.Args() {

		// Load the data from file.
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(filename + ": unable to read file")
			continue
		}

		// Declare the hashes.
		md5Hash := md5.New()
		sha256Hash := sha256.New()

		// Calculate the hashes on demand.
		if md5Flag {
			md5Hash.Write(data)
		}
		if sha256Flag {
			sha256Hash.Write(data)
		}

		// Print the output.
		fmt.Println("\nfile:   " + filename)
		if md5Flag {
			fmt.Println("md5:    " + hex.EncodeToString(md5Hash.Sum(nil)))
		}
		if sha256Flag {
			fmt.Println("sha256: " + hex.EncodeToString(sha256Hash.Sum(nil)))
		}
	}
}
