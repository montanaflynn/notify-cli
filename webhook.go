package notifier

import (
	"encoding/json"
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
		fmt.Printf("Notify: Sending %s Request to %s with body %s\n", w.method, w.url, w.body)
	}

	client := http.Client{}
	req, err := http.NewRequest(w.method, w.url, strings.NewReader(w.body))
	if err != nil {
		return err
	}

	for k, v := range w.headers {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if w.verbose || w.debug {
		fmt.Printf("Notify: Webhook returned status: %d\n", res.StatusCode)
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
		var headers map[string]string
		err := json.Unmarshal([]byte(headersRaw.(string)), &headers)
		if err != nil {
			return nil, err
		}
		notifier.headers = headers
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
