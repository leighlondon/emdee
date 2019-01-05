package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
	"text/tabwriter"
)

var commit = "latest"

const version = "0.9.0"

const usage = `usage: emdee [options] filename...

    Calculate message digests for input files.

options:
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

type flusher interface {
	Flush() error
}

func run(stdout io.Writer, stderr io.Writer, opts *options, names []string) int {
	if w, ok := stdout.(flusher); ok {
		defer w.Flush()
	}
	if w, ok := stderr.(flusher); ok {
		defer w.Flush()
	}

	if opts.version {
		fmt.Fprintf(stdout, "emdee v%s, commit %s\n", version, commit)
		return 0
	}
	if !opts.md5 && !opts.sha1 {
		// default is sha256
		opts.sha256 = true
	}
	if len(names) == 0 {
		fmt.Fprintf(stderr, "no filename provided\n")
		return 2
	}

	for _, fn := range names {
		f, err := os.Open(fn)
		if err != nil {
			fmt.Fprintf(stderr, fn+": unable to read file\n")
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
			fmt.Fprintf(stderr, "%s\n", err)
			return 1
		}

		d := hex.EncodeToString(h.Sum(nil))
		fmt.Fprintf(stdout, "%s\t%s\n", fn, d)
	}

	return 0
}

func main() {
	stdout := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', 0)
	stderr := tabwriter.NewWriter(os.Stderr, 0, 0, 8, ' ', 0)

	opts := options{}
	flag.BoolVar(&opts.md5, "md5", false, "Calculate the MD5 hash.")
	flag.BoolVar(&opts.sha256, "sha256", false, "Calculate the SHA256 hash.")
	flag.BoolVar(&opts.sha1, "sha1", false, "Calculate the SHA1 hash.")
	flag.BoolVar(&opts.version, "v", false, "Show the version number.")
	flag.Usage = func() { fmt.Fprintf(os.Stdout, usage) }
	flag.Parse()

	os.Exit(run(stdout, stderr, &opts, flag.Args()))
}
