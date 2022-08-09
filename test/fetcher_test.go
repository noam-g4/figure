package test

import (
	"testing"

	"github.com/noam-g4/figure/fetcher"
)

func TestReadFileSuccess(t *testing.T) {
	err, bytes := fetcher.ReadFile("./resource/test-config.yml")
	if err != nil && bytes == nil {
		t.Fail()
	}
}

func TestReadFileFail(t *testing.T) {
	err, _ := fetcher.ReadFile("./resource/not-exsits.yml")
	if err == nil {
		t.Fail()
	}
}
