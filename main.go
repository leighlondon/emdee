package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"hash"
	"io"
	"log"
	"os"
)

var commit = "latest"

const version = "0.7.0"

const usage = "emdee " + version + `

    Calculate message digests for input files.

Usage:
    emdee [options] filename...

Options:
    -h      Show this screen.
    -md5    Calculate the MD5 hash.
    -sha256 Calculate the SHA256 hash. (default)
    -sha1   Calculate the SHA1 hash.
    -v      Show the version number.
`

type options struct {
	md5     bool
	sha256  bool
	sha1    bool
	version bool
}

func run(logger *log.Logger, opts *options, names []string) int {

	if opts.version {
		logger.Printf("emdee v%s, commit %s\n", version, commit)
		return 0
	}

	if !opts.md5 && !opts.sha1 {
		// default is sha256
		opts.sha256 = true
	}

	if len(names) == 0 {
		logger.Println("no filename provided")
		return 2
	}

	for _, fn := range names {

		f, err := os.Open(fn)
		if err != nil {
			logger.Println(fn + ": unable to read file")
			continue
		}
		defer f.Close()

		rdr := bufio.NewReader(f)
		var h hash.Hash
		switch {
		case opts.md5:
			h = md5.New()
		case opts.sha1:
			h = sha1.New()
		case opts.sha256:
			h = sha256.New()
		}

		_, err = io.Copy(h, rdr)
		if err != nil {
			logger.Printf("%s", err)
			return 1
		}

		d := hex.EncodeToString(h.Sum(nil))
		logger.Printf("%s\t%s\n", fn, d)
	}
	return 0
}

func main() {
	logger := log.New(os.Stdout, "", 0)

	opts := options{}
	flag.BoolVar(&opts.md5, "md5", false, "Calculate the MD5 hash.")
	flag.BoolVar(&opts.sha256, "sha256", false, "Calculate the SHA256 hash.")
	flag.BoolVar(&opts.sha1, "sha1", false, "Calculate the SHA1 hash.")
	flag.BoolVar(&opts.version, "v", false, "Show the version number.")
	flag.Usage = func() { logger.Printf(usage) }
	flag.Parse()

	os.Exit(run(logger, &opts, flag.Args()))
}
