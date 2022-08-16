package test

import (
	"testing"

	"github.com/noam-g4/figure/env"
	"github.com/noam-g4/figure/fetcher"
	"github.com/noam-g4/figure/modifier"
	"github.com/noam-g4/figure/parser"
)

func TestTracePath(t *testing.T) {
	_, d := fetcher.ReadFile("./resource/test-config.yml")
	_, m := parser.ParseToMap(d)

	p1 := modifier.TracePath("child", m, modifier.Path{})
	if p1[0] != "parent" && p1[1] != "child" {
		t.Error(p1)
	}

	p2 := modifier.TracePath("env", m, modifier.Path{})
	if p2[0] != "env" {
		t.Error(p2)
	}

	p3 := modifier.TracePath("notExist", m, modifier.Path{})
	if len(p3) > 0 {
		t.Error(p3)
	}
}

func TestGetModifier(t *testing.T) {
	_, d := fetcher.ReadFile("./resource/test-config.yml")
	_, m := parser.ParseToMap(d)

	if m1 := modifier.GetModifier(env.Var{
		Name:  "env",
		Value: "test",
	}, m); len(m1.Path) != 1 && m1.Path[0] != "env" {
		t.Fail()
	}

	if m2 := modifier.GetModifier(env.Var{
		Name:  "child",
		Value: "6",
	}, m); len(m2.Path) != 2 {
		t.Fail()
	}

	if m3 := modifier.GetModifier(env.Var{
		Name:  "none",
		Value: "null",
	}, m); len(m3.Path) != 0 {
		t.Fail()
	}
}
