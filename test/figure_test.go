package test

import (
	"os"
	"strconv"
	"testing"

	"github.com/noam-g4/figure"
	"github.com/noam-g4/figure/replacer"
)

const path = "./resource/test-config.yml"
const invalid = "invalidpath"

type Cnfg struct {
	Env       string `yaml:"env"`
	WriteMode bool   `yaml:"writeMode"`
	Retries   int    `yaml:"retries"`
}

func TestLoadConfigWithoutReplacers(t *testing.T) {

	err, cSucc := figure.LoadConfigWithoutReplacers[Cnfg](path)

	if err != nil || cSucc.Env != "test" || cSucc.Retries != 5 || cSucc.WriteMode {
		t.Fail()
	}

	err, _ = figure.LoadConfigWithoutReplacers[Cnfg](invalid)
	if err == nil {
		t.Fail()
	}

}

func TestLoadConfig(t *testing.T) {
	os.Setenv("FIG_WRITEMODE", "true")
	os.Setenv("FIG_RETRIES", "9")

	replacers := []replacer.Replacer[Cnfg]{
		{
			Env: "FIG_WRITEMODE",
			Setter: func(c Cnfg, v string) Cnfg {
				if v != "true" {
					return c
				}
				newConf := c
				newConf.WriteMode = true
				return newConf
			},
		},
		{
			Env: "FIG_RETRIES",
			Setter: func(c Cnfg, v string) Cnfg {
				i, err := strconv.Atoi(v)
				if err != nil {
					return c
				}
				newConf := c
				newConf.Retries = i
				return newConf
			},
		},
	}

	err, cSucc := figure.LoadConfig(path, replacers)

	if err != nil || cSucc.Env != "test" || cSucc.Retries != 9 || !cSucc.WriteMode {
		t.Fail()
	}

	err, _ = figure.LoadConfig(invalid, replacers)

	if err == nil {
		t.Fail()
	}

}
