package traininfo

import (
	"errors"
	"net/http"
	"traininfobot/models"
	"traininfobot/util"
)

type Traininfo interface {
	TraininfoFunc() error
}

type traininfo struct {
	slackUtil util.SlackUtil
	httpUtil  util.HTTPUtil
}

func NewTraininfo(slackutil util.SlackUtil, httputil util.HTTPUtil) Traininfo {
	return &traininfo{slackutil, httputil}
}

type TrainDeley []struct {
	Name          string `json:"name"`
	Company       string `json:"company"`
	LastupdateGmt int    `json:"lastupdate_gmt"`
	Source        string `json:"source"`
}

func (t traininfo) TraininfoFunc() error {
	var responseJSON TrainDeley

	statuscode, err := t.httpUtil.GET("", "", &responseJSON)
	if err != nil {
		return err
	}
	if statuscode != http.StatusOK {
		return errors.New("API応答が正常ではありません")
	}
	var message models.Message
	var deleycount = 0
	var trains []models.TrainCompany
	trains = append(trains, models.TrainCompany{
		CompanyName: "名古屋鉄道",
		URL:         "https://top.meitetsu.co.jp/em/",
	})
	trains = append(trains, models.TrainCompany{
		CompanyName: "名古屋市営地下鉄",
		URL:         "https://www.kotsu.city.nagoya.jp/jp/pc/emergency/index.html#HIG",
	})
	for _, traindeley := range responseJSON {
		if traindeley.Company == "名古屋鉄道" && (traindeley.Name == "名鉄線" || traindeley.Name == "名古屋本線") {
			train := trains[0]
			message = models.Message{
				Delay: true,
				Text:  traindeley.Name + "が遅れてるよ！",
			}
			err = t.slackUtil.PostSlackMessage(train, message)
			deleycount++
		}
		if traindeley.Company == "名古屋市交通局" {
			train := trains[1]
			message = models.Message{
				Delay: true,
				Text:  traindeley.Name + "が遅れてるよ！",
			}
			err = t.slackUtil.PostSlackMessage(train, message)
			deleycount++
		}
	}
	if deleycount == 0 {
		for _, train := range trains {
			message = models.Message{
				Delay: false,
				Text:  "通常通りだよ！",
			}
			err = t.slackUtil.PostSlackMessage(train, message)
		}
	}
	if err != nil {
		return err
	}
	return nil
}
