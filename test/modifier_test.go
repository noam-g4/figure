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
