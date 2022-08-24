package test

import (
	"os"
	"testing"

	"github.com/noam-g4/figure/v2"
)

type Conf struct {
	Env       string `yaml:"env"`
	WriteMode bool   `yaml:"writeMode"`
	Others    struct {
		Retries int      `yaml:"retries"`
		Options []string `yaml:"options"`
	} `yaml:"others"`
}

func TestLoadConfig(t *testing.T) {
	os.Setenv("TST_ENV", "modified")
	os.Setenv("TST_WRITE_MODE", "TRUE")
	os.Setenv("TST_RETRIES", "5")
	os.Setenv("TST_OPTIONS", "[a,b,c]")

	err, conf := figure.LoadConfig[Conf](figure.Settings{
		FilePath:   "./resource/modify-test.yml",
		Prefix:     "TST_",
		Convention: figure.Camel,
		Separator:  "_",
	})

	if err != nil {
		t.Fail()
	}

	if conf.Env != "modified" {
		t.Error(conf.Env)
	}

	if !conf.WriteMode {
		t.Error(conf.WriteMode)
	}

	if conf.Others.Retries != 5 {
		t.Error(conf.Others.Retries)
	}

	if len(conf.Others.Options) != 3 && conf.Others.Options[1] != "b" {
		t.Error(conf.Others.Options)
	}
}
