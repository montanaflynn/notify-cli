package main

import (
	"bufio"
	"log"
	"os"

	notifier "github.com/montanaflynn/notify-cli"
)

func main() {
	opts, err := notifier.Usage()
	if err != nil {
		log.Fatal(err)
	}

	n, err := notifier.New(*opts)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	err = notifier.CommandPipe(r, w, n)
	if err != nil {
		log.Fatal(err)
	}
}
