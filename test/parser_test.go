package test

import (
	"testing"

	"github.com/noam-g4/figure/v2"
	"github.com/noam-g4/figure/v2/env"
	"github.com/noam-g4/figure/v2/fetcher"
	"github.com/noam-g4/figure/v2/modifier"
	"github.com/noam-g4/figure/v2/parser"
)

const realFile = "./resource/test-config.yml"

type Config struct {
	One   int    `yaml:"one"`
	Two   string `yaml:"two"`
	Three struct {
		Four struct {
			Five []int `yaml:"five"`
		} `yaml:"four"`
	} `yaml:"three"`
	Nine bool `yaml:"nine"`
	Ten  struct {
		Eleven int `yaml:"eleven"`
	} `yaml:"ten"`
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

	camel := parser.TransformName(figure.Camel, "-", v)
	snake := parser.TransformName(figure.Snake, "-", v)
	caps := parser.TransformName(figure.Caps, "-", v)
	snakeCaps := parser.TransformName(figure.SnakeCaps, "-", v)
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
		m["ten"].(map[string]interface{})["eleven"] != 12 {
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
	if err != nil || conf.One != 1 || conf.Two != "two" || len(conf.Three.Four.Five) != 3 {
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

// testing with modification of the origianl
func TestSerializeYamlMap(t *testing.T) {
	_, data := fetcher.ReadFile(realFile)
	_, m := parser.ParseToMap(data)

	v := env.Var{Name: "eleven", Value: "20"}
	_, i := parser.CastIntValue(v.Value.(string))
	v.Value = i

	mod := modifier.Modify(v, m)
	err, bts := parser.SerializeYamlMap(mod)
	if err != nil || bts == nil {
		t.Fail()
	}

	err, conf := parser.Parse[Config](bts)
	if err != nil || conf.Ten.Eleven != 20 {
		t.Fail()
	}
}
