package main

import (
	"bufio"
	"log"
	"os"

	"github.com/montanaflynn/notify-cli"
)

var (
	version = "unknown"
	commit  = "unknown"
)

func main() {
	opts, err := notify.Usage(version, commit)
	if err != nil {
		log.Fatal(err)
	}

	n, err := notify.New(*opts)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	err = notify.CommandPipe(r, w, n)
	if err != nil {
		log.Fatal(err)
	}
}
