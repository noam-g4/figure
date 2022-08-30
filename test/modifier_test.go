package test

import (
	"testing"

	"github.com/noam-g4/figure/v2/env"
	"github.com/noam-g4/figure/v2/fetcher"
	"github.com/noam-g4/figure/v2/modifier"
	"github.com/noam-g4/figure/v2/parser"
)

func TestModify(t *testing.T) {
	_, d := fetcher.ReadFile("./resource/test-config.yml")
	_, m := parser.ParseToMap(d)

	v := env.Var{
		Name:  "one",
		Value: "15",
	}

	if res := modifier.Modify(v, m, "__"); res["one"] != "15" && m["one"] != 1 {
		t.Fail()
	}

	v2 := env.Var{
		Name:  "three__four__five",
		Value: 5,
	}
	r := modifier.Modify(v2, m, "__")
	val := r["three"].(map[interface{}]interface{})["four"].(map[interface{}]interface{})["five"]
	if val != 5 {
		t.Error(val)
	}
}

func TestUpdateMapWithEnvs(t *testing.T) {
	envs := []env.Var{
		{
			Name:  "env",
			Value: "modified",
		},
		{
			Name:  "writeMode",
			Value: "true",
		},
		{
			Name:  "others__retries",
			Value: "5",
		},
		{
			Name:  "others__options",
			Value: "[a, b, c]",
		},
	}

	_, data := fetcher.ReadFile("./resource/modify-test.yml")
	_, mp := parser.ParseToMap(data)

	nMap := modifier.UpdateMapWithEnvs(envs, mp, "__")
	if nMap["env"] != "modified" {
		t.Error(nMap)
	}
	if !nMap["writeMode"].(bool) {
		t.Error(nMap)
	}
	others := nMap["others"].(map[interface{}]interface{})
	if others["retries"] != 5 {
		t.Error(nMap)
	}
	options := others["options"].([]interface{})
	if len(options) != 3 && options[2] != "b" {
		t.Fail()
	}

}

func TestValueOrDefaultMapSep(t *testing.T) {
	d := modifier.ValueOrDefault("", "__")
	v := modifier.ValueOrDefault(":", "_")
	if d != "__" {
		t.Error(d)
	}
	if v != ":" {
		t.Error(v)
	}
}
