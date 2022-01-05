package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hguerra/envsubst/v2"
	"log"
	"os"
)

var (
	noEmpty = flag.Bool("no-empty", false, "")
)

var usage = `Usage: envsubst [options...] <input>
Options:
  -no-empty  Fail if a variable is set but empty.
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()

	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	for stdin.Scan() {
		line, err := envsubst.EvalEnv(stdin.Text(), *noEmpty)
		if err != nil {
			log.Fatalf("Error while envsubst: %v", err)
		}
		_, err = fmt.Fprintln(stdout, line)
		if err != nil {
			log.Fatalf("Error while writing to stdout: %v", err)
		}
		stdout.Flush()
	}
}
