package notifier

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/docopt/docopt-go"
)

var usageString = `Notify
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
  --headers=<headers>  Request headers [optional, default: '[]', syntax: '[api-key=token, user-agent=notify-cli]']`

// Usage parses command line command and flags
func Usage() (*docopt.Opts, error) {
	options, err := docopt.ParseDoc(usageString)
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