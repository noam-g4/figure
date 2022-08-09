package test

import (
	"os"
	"strconv"
	"testing"

	"github.com/noam-g4/figure/fetcher"
	"github.com/noam-g4/figure/parser"
	"github.com/noam-g4/figure/replacer"
	r "github.com/noam-g4/figure/replacer"
)

type config struct {
	Env       string `yaml:"env"`
	WriteMode bool   `yaml:"writeMode"`
	Retries   int    `yaml:"retries"`
}

func TestGetEnv(t *testing.T) {
	os.Setenv("FIGURE_TEST", "OK")
	ok, val := replacer.GetEnv("FIGURE_TEST")
	if !ok || val != "OK" {
		t.Fail()
	}

	ok, val = r.GetEnv("FIGURE_NOT")
	if ok && val != "" {
		t.Fail()
	}
}

func TestReplaceConfigWithEnv(t *testing.T) {
	rs := []r.Replacer[config]{
		{
			Env: "FIGURE_WRITEMODE",
			Setter: func(c config, v string) config {
				if v != "true" {
					return c
				}
				nc := c
				nc.WriteMode = true
				return nc
			},
		},
		{
			Env: "FIGURE_RETRIES",
			Setter: func(c config, v string) config {
				i, err := strconv.Atoi(v)
				if err != nil {
					return c
				}
				nc := c
				nc.Retries = i
				return nc
			},
		},
	}

	_, data := fetcher.ReadFile("./resource/test-config.yml")
	_, conf := parser.Parse[config](data)
	os.Setenv("FIGURE_WRITEMODE", "true")
	os.Setenv("FIGURE_RETRIES", "7")

	newConf := r.ReplaceConfigWithEnv(conf, rs)
	if newConf.Retries != 7 || !newConf.WriteMode {
		t.Fail()
	}

}
