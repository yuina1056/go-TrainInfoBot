package traininfo

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestTrainInfo(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't .env load.")
	}
	err = traininfo()
	if err != nil {
		t.Fatal(err)
	}
}
