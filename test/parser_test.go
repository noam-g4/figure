package test

import (
	"testing"

	"github.com/noam-g4/figure/env"
	"github.com/noam-g4/figure/fetcher"
	"github.com/noam-g4/figure/parser"
)

const realFile = "./resource/test-config.yml"

type Config struct {
	Env string `yaml:"env"`
}

func TestStripPrefix(t *testing.T) {
	v := env.Var{
		Name:  "TST_ONE",
		Value: "one",
	}

	res := parser.StripPrefix("TST_", v)
	if res.Name != "ONE" {
		t.Fail()
	}
}

func TestTransformName(t *testing.T) {
	v := env.Var{
		Name:  "TEST-this",
		Value: "null",
	}

	camel := parser.TransformName(parser.Camel, "-", v)
	snake := parser.TransformName(parser.Snake, "-", v)
	caps := parser.TransformName(parser.Caps, "-", v)
	snakeCaps := parser.TransformName(parser.SnakeCaps, "-", v)
	undifiend := parser.TransformName(4, "-", v)

	if camel.Name != "testThis" {
		t.Fail()
	}
	if snake.Name != "test_this" {
		t.Fail()
	}
	if caps.Name != "TEST-THIS" {
		t.Fail()
	}
	if snakeCaps.Name != "TEST_THIS" {
		t.Fail()
	}
	if undifiend.Name != "TEST-this" {
		t.Fail()
	}
}

func TestParseToMapSuccess(t *testing.T) {
	_, data := fetcher.ReadFile(realFile)
	err, m := parser.ParseToMap(data)
	if err != nil &&
		m["parent"].(map[string]interface{})["child"] != 3 {
		t.Fail()
	}
}

func TestParseToMapFail(t *testing.T) {
	data := []byte("not a valid yaml")
	err, _ := parser.ParseToMap(data)
	if err == nil {
		t.Fail()
	}
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
