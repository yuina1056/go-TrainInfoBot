package util

import (
	"os"
	"testing"
	"traininfobot/models"

	"github.com/joho/godotenv"
)

func TestPostSlackBotnormal(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal("can't .env load.")
	}
	train := models.TrainCompany{
		CompanyName: "名古屋鉄道",
		URL:         "https://top.meitetsu.co.jp/em/",
	}
	message := models.Message{
		Delay: false,
		Text:  "通常通りだよ！",
	}
	slackutil := NewSlackUtil(os.Getenv("SLACK_WEBHOOK_URL"), os.Getenv("SLACK_MENTION"))
	err = slackutil.PostSlackMessage(train, message)
	if err != nil {
		t.Fatal("SlackPost Error.")
	}
}

func TestPostSlackBotdelay(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal("can't .env load.")
	}
	train := models.TrainCompany{
		CompanyName: "名古屋鉄道",
		URL:         "https://top.meitetsu.co.jp/em/",
	}
	message := models.Message{
		Delay: true,
		Text:  "遅れてるよ！",
	}
	slackutil := NewSlackUtil(os.Getenv("SLACK_WEBHOOK_URL"), os.Getenv("SLACK_MENTION"))
	err = slackutil.PostSlackMessage(train, message)
	if err != nil {
		t.Fatal("SlackPost Error.")
	}
}
