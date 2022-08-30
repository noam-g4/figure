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

	camel := parser.TransformName(figure.Camel, "-", "__", v)
	snake := parser.TransformName(figure.Snake, "-", "__", v)
	caps := parser.TransformName(figure.Caps, "-", "__", v)
	snakeCaps := parser.TransformName(figure.SnakeCaps, "-", "__", v)
	undifiend := parser.TransformName(4, "-", "__", v)

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

	v := env.Var{Name: "ten__eleven", Value: "20"}
	_, i := parser.CastIntValue(v.Value.(string))
	v.Value = i

	mod := modifier.Modify(v, m, "__")
	err, bts := parser.SerializeYamlMap(mod)
	if err != nil || bts == nil {
		t.Error(err)
	}

	err, conf := parser.Parse[Config](bts)
	if err != nil || conf.Ten.Eleven != 20 {
		t.Error(mod)
	}
}

func TestSmartSplit(t *testing.T) {
	s1 := "options__write_mode"
	s2 := "mod"
	s3 := "num:of:tries"
	s4 := ""
	s5 := "main_service__options__write_mode"
	s6 := "write_mode"

	r1 := parser.SmartSplit(s1, "_", "__")
	if len(r1) != 2 || r1[1] != "mode" {
		t.Error(r1)
	}

	if parser.SmartSplit(s2, "/", ":")[0] != "mod" {
		t.Error(s2)
	}

	r2 := parser.SmartSplit(s3, ":", "__")
	if len(s3) != 3 && r2[1] != "of" {
		t.Error(r2)
	}

	if len(parser.SmartSplit(s4, "", "")) != 0 {
		t.Error("not empty")
	}

	r3 := parser.TransformName(
		figure.Camel,
		"_",
		"__",
		env.Var{
			Name:  s5,
			Value: 5,
		},
	)

	if r3.Name != "mainService__options__writeMode" {
		t.Error(r3.Name)
	}

	r4 := parser.SmartSplit(s6, "_", "__")
	if len(r4) != 2 && r4[0] != "write" && r4[1] != "mode" {
		t.Error(r4)
	}

}
