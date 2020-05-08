package notifier

import (
	"bufio"
	"io"

	"github.com/docopt/docopt-go"
)

// Notifier is the interface for sending a notification
type Notifier interface {
	Notify() error
}

// New creates a new Notifier from the command and flags
func New(opts docopt.Opts) (Notifier, error) {
	isSlack, err := opts.Bool("slack")
	if err != nil {
		return nil, err
	}

	isWebhook, err := opts.Bool("webhook")
	if err != nil {
		return nil, err
	}

	switch {
	case isSlack:
		return newSlackNotifier(opts)
	case isWebhook:
		return newWebhookNotifier(opts)
	}

	return &bellNotifier{}, nil
}

// CommandPipe pipes input from a reader to a writer
// and then triggers a notification when done
func CommandPipe(r *bufio.Reader, w *bufio.Writer, n Notifier) error {
	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			return n.Notify()
		} else if err != nil {
			return err
		}
		w.WriteByte(b)
		w.Flush()
	}
}
