package test

import (
	"testing"

	"github.com/noam-g4/figure/env"
	"github.com/noam-g4/figure/fetcher"
	"github.com/noam-g4/figure/modifier"
	"github.com/noam-g4/figure/parser"
)

func TestModify(t *testing.T) {
	_, d := fetcher.ReadFile("./resource/test-config.yml")
	_, m := parser.ParseToMap(d)

	v := env.Var{
		Name:  "one",
		Value: "15",
	}

	if res := modifier.Modify(v, m); res["one"] != "15" && m["one"] != 1 {
		t.Fail()
	}

	v2 := env.Var{
		Name:  "five",
		Value: "mod",
	}
	r := modifier.Modify(v2, m)
	val := r["three"].(map[interface{}]interface{})["four"].(map[interface{}]interface{})["five"]
	if val != "mod" {
		t.Fail()
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
			Name:  "retries",
			Value: "5",
		},
		{
			Name:  "options",
			Value: "[a, b, c]",
		},
	}

	_, data := fetcher.ReadFile("./resource/modify-test.yml")
	_, mp := parser.ParseToMap(data)

	nMap := modifier.UpdateMapWithEnvs(envs, mp)
	if nMap["env"] != "modified" {
		t.Fail()
	}
	if !nMap["writeMode"].(bool) {
		t.Fail()
	}
	others := nMap["others"].(map[interface{}]interface{})
	if others["retries"] != 5 {
		t.Fail()
	}
	options := others["options"].([]interface{})
	if len(options) != 3 && options[2] != "b" {
		t.Fail()
	}

}
