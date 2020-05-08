package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/docopt/docopt-go"
)

type notifier interface {
	Notify() error
}

func sendNotification(options docopt.Opts) error {
	isSlack, err := options.Bool("slack")
	if err != nil {
		return err
	}

	isWebhook, err := options.Bool("webhook")
	if err != nil {
		return err
	}

	switch {
	case isSlack:
		slack, err := newSlackNotifier(options)
		if err != nil {
			return err
		}
		return slack.Notify()
	case isWebhook:
		webhook, err := newWebhookNotifier(options)
		if err != nil {
			return err
		}
		return webhook.Notify()
	}

	bell := bellNotifier{}
	return bell.Notify()
}

func notify(options docopt.Opts) error {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			return sendNotification(options)
		} else if err != nil {
			return err
		}
		w.WriteByte(b)
		w.Flush()
	}
}

func main() {
	usage := `Notify CLI
Usage:
  notify-cli
  notify-cli slack --token=<token> --channel=<channel> [--verbose] [--debug]
  notify-cli webhook --url=<url> [--method=<method>] [--body=<body>] [--headers=<headers>] [--verbose] [--debug]
  notify-cli -h | --help
  notify-cli --version

Flags:
  -h --help            Show this screen.
  --version            Show version.
  --verbose            Show verbose output. [optional, default: false, syntax: BOOL]
  --debug              Show debug output. [optional, default: false, syntax: BOOL]
  --token=<token>      Slack token [required, syntax: 'STRING']
  --channel=<channel>  Slack channel [required, syntax: 'STRING']
  --url=<url>          Request URL [required, syntax: 'STRING']
  --method=<method>    Request method [optional, default: "GET", syntax: 'STRING']
  --body=<body>        Request body [optional, default: '', syntax: 'STRING']
  --headers=<headers>  Request headers [optional, default: '[]', syntax: '[api-key=token, user-agent=notify-cli]']`

	options, err := docopt.ParseDoc(usage)
	if err != nil {
		log.Fatal(err)
	}

	debugInterface, _ := options["--debug"]
	if debugInterface != nil {
		debug := debugInterface.(bool)
		if debug {
			fmt.Println("Options:")
			spew.Dump(options)
		}
	}

	err = notify(options)
	if err != nil {
		log.Fatal(err)
	}
}
