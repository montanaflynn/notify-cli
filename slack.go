package notifier

import (
	"fmt"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/slack-go/slack"
)

type slackNotifier struct {
	token   string
	channel string
	message string
	debug   bool
	verbose bool
}

func (s *slackNotifier) Notify() error {
	api := slack.New(s.token, slack.OptionDebug(s.debug))
	channels, err := api.GetChannels(true)
	if err != nil {
		return err
	}
	var channelID string
	for _, channel := range channels {
		if strings.ToLower(channel.Name) == strings.ToLower(s.channel) {
			channelID = channel.ID
		}
	}

	if channelID == "" {
		return fmt.Errorf("Notify: Could not find channel: %s", s.channel)
	}

	if s.verbose {
		fmt.Printf("Notify: Sending message %s to slack\n", s.message)
	}

	text := s.message
	messageText := slack.NewTextBlockObject("mrkdwn", text, false, false)
	messageSectionBlock := slack.NewSectionBlock(messageText, nil, nil)
	fallbackOpt := slack.MsgOptionText(text, false)
	blockOpt := slack.MsgOptionBlocks(messageSectionBlock)
	msg := slack.MsgOptionCompose(fallbackOpt, blockOpt)
	_, _, err = api.PostMessage(channelID, msg)
	if err != nil {
		return err
	}

	if s.verbose {
		fmt.Println("Notify: Sent message to slack without an error")
	}

	return nil
}

func newSlackNotifier(options docopt.Opts) (*slackNotifier, error) {
	token, err := options.String("--token")
	if err != nil {
		return nil, err
	}

	channel, err := options.String("--channel")
	if err != nil {
		return nil, err
	}

	notifier := &slackNotifier{
		token:   token,
		channel: channel,
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
