package notifier

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/docopt/docopt-go"
)

var (
	version = "unknown"
	commit  = "unknown"
)

var usageString = strings.TrimSpace(`
Notify

Usage:
  notify
  notify slack --token=<token> --channel=<channel> [--verbose] [--debug]
  notify webhook --url=<url> [--method=<method>] [--body=<body>] [--headers=<headers>] [--verbose] [--debug]
  notify -h | --help
  notify --version

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
  --headers=<headers>  Request headers [optional, default: '{}', syntax: '{"key":"val", "foo":"bar"}']

Examples:
  echo "hello world" | notify
  echo "hello world" | notify slack --token='xxxx-xxxxxxxxx-xxxx' --channel='#notifications'
  echo "hello world" | notify webhook --url=http://requestbin.net/r/1897yhr1 --headers='{"key":"val", "foo":"bar"}'
`)

// Usage parses command line command and flags
func Usage() (*docopt.Opts, error) {
	version := fmt.Sprintf("Version: %s\nCommit: %s", version, commit)
	options, err := docopt.ParseArgs(usageString, os.Args[1:], version)
	if err != nil {
		return nil, err
	}

	debugInterface, _ := options["--debug"]
	if debugInterface != nil {
		debug := debugInterface.(bool)
		if debug {
			fmt.Println("Options:")
			spew.Dump(options)
		}
	}

	return &options, nil
}
