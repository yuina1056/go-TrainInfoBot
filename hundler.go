package main

import (
	"fmt"
	"os"
	"traininfobot/traininfo"
	"traininfobot/util"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func hundler() (MyResponse, error) {
	slackutil := util.NewSlackUtil(os.Getenv("SLACK_WEBHOOK_URL"), os.Getenv("SLACK_MENTION"))
	httputil := util.NewHTTPUtil(os.Getenv("TRAININFO_URL"))
	traininfo := traininfo.NewTraininfo(slackutil, httputil)
	err := traininfo.TraininfoFunc()
	if err != nil {
		return MyResponse{Message: fmt.Sprintf("Hello!!")}, err
	}
	return MyResponse{Message: fmt.Sprintf("Hello!!")}, nil
}

func main() {
	lambda.Start(hundler)
}
