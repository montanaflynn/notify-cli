# Notify CLI

Get notified when a command finishes.

### Why?

Instead of constantly checking a long running command you can get notified when it finishes.

### How?

Install `notify-cli` pre-compiled from `gobinaries` or from source if you have `go` installed:

```sh
# Using gobinaries
curl -sf https://gobinaries.com/montanaflynn/notify-cli | sh

# From source
go get github.com/montanaflynn/notify-cli
```

Then add `| notify-cli` to the end of a command. By default it makes a bell sound when the command finishes but can be configured to send a slack message or a webhook instead. Here's the usage from `notify-cli --help`:

```
Notify CLI

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
  --headers=<headers>  Request headers [optional, default: '[]', syntax: '[api-key=token, user-agent=notify-cli]']
```