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

var commit = "latest"

const version = "0.6.0"

const usage = "emdee " + version + `

    Calculate message digests for input files.

Usage:
    emdee [options] filename...

Options:
    -h      Show this screen.
    -md5    Calculate the MD5 hash.
    -sha256 Calculate the SHA256 hash.
    -sha1   Calculate the SHA1 hash.
    -v      Show the version number.
`

type options struct {
	md5     bool
	sha256  bool
	sha1    bool
	version bool
}

func main() {

	opts := options{}
	flag.BoolVar(&opts.md5, "md5", false, "Calculate the MD5 hash.")
	flag.BoolVar(&opts.sha256, "sha256", false, "Calculate the SHA256 hash.")
	flag.BoolVar(&opts.sha1, "sha1", false, "Calculate the SHA1 hash.")
	flag.BoolVar(&opts.version, "v", false, "Show the version number.")
	flag.Usage = func() { fmt.Printf(usage) }
	flag.Parse()

	// Check for the easy flag.
	if opts.version {
		fmt.Printf("emdee v%s, commit %s\n", version, commit)
		return
	}

	// If no flags were set, run everything.
	if flag.NFlag() == 0 {
		opts.md5 = true
		opts.sha1 = true
		opts.sha256 = true
	}

	if flag.NArg() == 0 {
		fmt.Println("no filename provided")
		return
	}

	fmt.Println(version)

	for _, fn := range flag.Args() {

		f, err := os.Open(fn)
		if err != nil {
			fmt.Println("\n" + fn + ": unable to read file")
			continue
		}
		defer f.Close()

		rdr := bufio.NewReader(f)

		// Declare the hashes.
		h0 := md5.New()
		h1 := sha256.New()
		h2 := sha1.New()

		// Create a MultiWriter of all of the hashes.
		all := io.MultiWriter(h0, h1, h2)

		// Copy to all at once.
		_, err = io.Copy(all, rdr)
		if err != nil {
			break
		}

		// Print the output.
		fmt.Println("\nfile:   " + fn)
		if opts.md5 {
			fmt.Println("md5:    " + hex.EncodeToString(h0.Sum(nil)))
		}
		if opts.sha1 {
			fmt.Println("sha1:   " + hex.EncodeToString(h2.Sum(nil)))
		}
		if opts.sha256 {
			fmt.Println("sha256: " + hex.EncodeToString(h1.Sum(nil)))
		}
	}
}
