package test

import (
	"testing"

	"github.com/noam-g4/figure/env"
	"github.com/noam-g4/figure/fetcher"
	"github.com/noam-g4/figure/modifier"
	"github.com/noam-g4/figure/parser"
)

func TestFindValue(t *testing.T) {
	_, d := fetcher.ReadFile("./resource/test-config.yml")
	_, m := parser.ParseToMap(d)

	p1 := modifier.FindValue("eleven", m)
	if *p1 != 12 {
		t.Fail()
	}

	p2 := modifier.FindValue("two", m)
	if *p2 != "two" {
		t.Fail()
	}

	p3 := modifier.FindValue("five", m)
	if p3 == nil {
		t.Fail()
	}

	p4 := modifier.FindValue("twelve", m)
	if p4 != nil {
		t.Fail()
	}
}

func TestGetModifier(t *testing.T) {
	_, d := fetcher.ReadFile("./resource/test-config.yml")
	_, m := parser.ParseToMap(d)

	if m1 := modifier.GetModifier(env.Var{
		Name:  "eleven",
		Value: "12",
	}, m); *m1.Accessor != 12 {
		t.Fail()
	}

	if m2 := modifier.GetModifier(env.Var{
		Name:  "two",
		Value: "three",
	}, m); *m2.Accessor != "two" {
		t.Fail()
	}
}
