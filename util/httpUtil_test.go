package util

import (
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

type TrainDeley []struct {
	Name          string `json:"name"`
	Company       string `json:"company"`
	LastupdateGmt int    `json:"lastupdate_gmt"`
	Source        string `json:"source"`
}

func TestGET(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal("can't .env load.")
	}

	var responseJSON TrainDeley

	HTTPUtil := NewHTTPUtil(os.Getenv("TRAININFO_URL"))
	statuscode, err := HTTPUtil.GET("", "", responseJSON)
	if err != nil {
		t.Fatal(err)
	}
	if statuscode != http.StatusOK {
		t.Fatal(statuscode)
	}
	t.Log(responseJSON)
}

func TestTraininfoGET(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal("can't .env load.")
	}

	var responseJSON TrainDeley

	HTTPUtil := NewHTTPUtil(os.Getenv("TRAININFO_URL"))
	statuscode, err := HTTPUtil.GET("", "", &responseJSON)
	if err != nil {
		t.Fatal(err)
	}
	if statuscode != http.StatusOK {
		t.Fatal(statuscode)
	}
	t.Log(responseJSON)
}
