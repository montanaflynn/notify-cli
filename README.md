# Notify CLI

Instead of constantly checking a long running command you can get notified when it finishes.

Here's an example that sends a notification to slack:

```sh
sleep 10 && echo "hello world" | notify slack --token='xxxx-xxxxxxxxx-xxxx' --channel='#notifications'
hello world
```

### Install

You can download and install the latest version of `notify` from [GitHub releases](https://github.com/montanaflynn/notify-cli/releases).

You can also install a pre-compiled `notify` from [gobinaries](https://gobinaries.com/) or from the latest source if you have `go` installed:

```sh
# Using gobinaries
curl -sf https://gobinaries.com/montanaflynn/notify-cli/cmd/notify | sh

# From source
go get github.com/montanaflynn/notify-cli/cmd/notify
```

### Usage

Once installed add `| notify` to the end of a command. By default it makes a bell sound when the command finishes but can be configured to send a slack message or a webhook instead. Here's the usage from `notify --help`:

```
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
```
