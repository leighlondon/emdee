package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

const usage = "emdee " + version + `

    Calculate message digests for input files.

Usage:
    emdee [options] filename..

Options:
    -h    Show this screen.
    -m    Calculate the MD5 hash.
    -s    Calculate the SHA256 hash.
    -s1   Calculate the SHA1 hash.
    -v    Show the version number.
`

func main() {

	// Flags for the program, declared here for scoping.
	md5Flag := flag.Bool("m", false, "Calculate the MD5 hash.")
	sha256Flag := flag.Bool("s", false, "Calculate the SHA256 hash.")
	sha1Flag := flag.Bool("s1", false, "Calculate the SHA1 hash.")
	versionFlag := flag.Bool("v", false, "Show the version number.")

	// Replace the usage screen.
	flag.Usage = func() {
		fmt.Printf(usage)
	}

	// Parse the flags.
	flag.Parse()

	// Check for the easy flag.
	if *versionFlag {
		fmt.Println(version)
		return
	}

	// If no flags were set, run everything.
	if flag.NFlag() == 0 {
		*md5Flag = true
		*sha1Flag = true
		*sha256Flag = true
	}

	// Need at least one filename to be provided.
	if flag.NArg() == 0 {
		fmt.Println("no filename provided")
		return
	}

	// Print the version string with the normal output.
	fmt.Println(version)

	for _, filename := range flag.Args() {

		// Load the data from file.
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("\n" + filename + ": unable to read file")
			continue
		}
		defer file.Close()

		// Use a buffered reader.
		reader := bufio.NewReader(file)

		// Declare the hashes.
		md5Hash := md5.New()
		sha256Hash := sha256.New()
		sha1Hash := sha1.New()

		// Create a MultiWriter of all of the hashes.
		all := io.MultiWriter(sha1Hash, sha256Hash, md5Hash)

		// Copy to all at once.
		_, err = io.Copy(all, reader)
		if err != nil {
			break
		}

		// Print the output.
		fmt.Println("\nfile:   " + filename)
		if *md5Flag {
			fmt.Println("md5:    " + hex.EncodeToString(md5Hash.Sum(nil)))
		}
		if *sha1Flag {
			fmt.Println("sha1:   " + hex.EncodeToString(sha1Hash.Sum(nil)))
		}
		if *sha256Flag {
			fmt.Println("sha256: " + hex.EncodeToString(sha256Hash.Sum(nil)))
		}
	}
}
