package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"io"
	"log"
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

func main() {
	log.SetFlags(0)

	opts := options{}
	flag.BoolVar(&opts.md5, "md5", false, "Calculate the MD5 hash.")
	flag.BoolVar(&opts.sha256, "sha256", false, "Calculate the SHA256 hash.")
	flag.BoolVar(&opts.sha1, "sha1", false, "Calculate the SHA1 hash.")
	flag.BoolVar(&opts.version, "v", false, "Show the version number.")
	flag.Usage = func() { log.Printf(usage) }
	flag.Parse()

	if opts.version {
		log.Printf("emdee v%s, commit %s\n", version, commit)
		return
	}

	if flag.NFlag() == 0 {
		// default is sha256
		opts.sha256 = true
	}

	if flag.NArg() == 0 {
		log.Println("no filename provided")
		return
	}

	for _, fn := range flag.Args() {

		f, err := os.Open(fn)
		if err != nil {
			log.Println("\n" + fn + ": unable to read file")
			continue
		}
		defer f.Close()

		rdr := bufio.NewReader(f)

		md := md5.New()
		s1 := sha1.New()
		s2 := sha256.New()

		all := io.MultiWriter(md, s1, s2)
		_, err = io.Copy(all, rdr)
		if err != nil {
			break
		}

		var d string
		switch {
		case opts.md5:
			d = hex.EncodeToString(md.Sum(nil))
		case opts.sha1:
			d = hex.EncodeToString(s1.Sum(nil))
		case opts.sha256:
			d = hex.EncodeToString(s2.Sum(nil))
		}
		log.Printf("%s\t%s\n", fn, d)
	}
}
