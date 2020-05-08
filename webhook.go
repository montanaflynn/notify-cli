package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/docopt/docopt-go"
)

type webhookNotifier struct {
	url     string
	method  string
	body    string
	headers map[string]string
	verbose bool
	debug   bool
}

func (w *webhookNotifier) Notify() error {
	if w.debug {
		fmt.Printf("Sending %s Request to %s with body %s\n", w.method, w.url, w.body)
	}

	client := http.Client{}
	req, err := http.NewRequest(w.method, w.url, strings.NewReader(w.body))
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if w.verbose || w.debug {
		fmt.Printf("Webhook returned status: %d\n", res.StatusCode)
	}
	return err
}

func newWebhookNotifier(options docopt.Opts) (*webhookNotifier, error) {
	url, err := options.String("--url")
	if err != nil {
		return nil, err
	}

	notifier := &webhookNotifier{
		url: url,
	}

	method, _ := options["--method"]
	if method != nil {
		notifier.method = method.(string)
	}

	body, _ := options["--body"]
	if body != nil {
		notifier.body = body.(string)
	}

	headersRaw, _ := options["--headers"]
	if headersRaw != nil {
		headersString := headersRaw.(string)
		headersSlice := strings.Split(headersString, ",")
		headersMap := make(map[string]string)
		for _, header := range headersSlice {
			headerSlice := strings.Split(header, "=")
			headersMap[strings.TrimSpace(headerSlice[0])] = strings.TrimSpace(headerSlice[1])
		}
		notifier.headers = headersMap
	}

	verbose, _ := options["--verbose"]
	if verbose != nil {
		notifier.verbose = verbose.(bool)
	}

	debug, _ := options["--debug"]
	if debug != nil {
		notifier.debug = debug.(bool)
	}

	return notifier, nil
}
