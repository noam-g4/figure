package test

import (
	"testing"

	"github.com/noam-g4/figure/fetcher"
	"github.com/noam-g4/figure/parser"
)

const realFile = "./resource/test-config.yml"

type Config struct {
	Env string `yaml:"env"`
}

func TestParseSuccess(t *testing.T) {
	_, data := fetcher.ReadFile(realFile)
	err, conf := parser.Parse[Config](data)
	if err != nil || conf.Env != "test" {
		t.Fail()
	}
}

func TestParseFail(t *testing.T) {
	data := []byte("not a valid yaml")
	err, _ := parser.Parse[Config](data)
	if err == nil {
		t.Fail()
	}
}

func wrapper[T interface{}]() T {
	_, data := fetcher.ReadFile(realFile)
	_, conf := parser.Parse[T](data)
	return conf
}

func TestGenericTypePropagation(t *testing.T) {
	conf := wrapper[Config]()
	if conf.Env != "test" {
		t.Fail()
	}
}
