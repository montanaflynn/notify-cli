package main

import (
	"bufio"
	"log"
	"os"

	"github.com/montanaflynn/notify-cli"
)

func main() {
	opts, err := notify.Usage()
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
