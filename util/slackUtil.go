package util

import (
	"errors"
	"traininfobot/models"

	"github.com/ashwanthkumar/slack-go-webhook"
)

// SlackUtil SlackUtilのインターフェース
type SlackUtil interface {
	PostSlackMessage(train models.TrainCompany, message models.Message) error
}

type slackUtil struct {
	WebhookURL string
	Mention    string
}

// NewSlackUtil SlackUtilを生成
func NewSlackUtil(webhookURL string, mention string) SlackUtil {
	return &slackUtil{webhookURL, mention}
}

// PostSlackMessage slackのWebHookに送信する
func (util slackUtil) PostSlackMessage(train models.TrainCompany, message models.Message) error {
	webhookURL := util.WebhookURL

	attachment1 := slack.Attachment{}
	attachment1.AddField(
		slack.Field{Title: train.CompanyName, Value: message.Text},
	)
	var text string
	if message.Delay {
		attachment1.AddAction(
			slack.Action{Type: "button", Text: "詳細", Url: train.URL, Style: "primary"},
		)
		text = util.Mention
	}
	payload := slack.Payload{
		Text:        text,
		Username:    "電車の情報",
		Channel:     "#_info_train",
		IconEmoji:   ":train2:",
		Attachments: []slack.Attachment{attachment1},
	}
	errs := slack.Send(webhookURL, "", payload)
	if len(errs) > 0 {
		var errString string
		for _, err := range errs {
			errString += string(err.Error()) + ","
		}
		err := errors.New(errString)
		return err
	}
	return nil
}
