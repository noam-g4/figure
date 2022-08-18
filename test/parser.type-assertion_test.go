package test

import (
	"testing"

	"github.com/noam-g4/figure/parser"
)

func TestCastBooleanValue(t *testing.T) {
	ok, res := parser.CastBooleanValue("  TruE ")
	if !ok || !res {
		t.Fail()
	}
	ok, res = parser.CastBooleanValue("FALSE ")
	if !ok || res {
		t.Fail()
	}
	ok, _ = parser.CastBooleanValue("not ok")
	if ok {
		t.Fail()
	}
}

func TestCastIntValue(t *testing.T) {
	ok, n := parser.CastIntValue(" 5")
	if n != 5 {
		t.Fail()
	}
	ok, _ = parser.CastIntValue("not a number")
	if ok {
		t.Fail()
	}
}

func TestCastFloatValue(t *testing.T) {
	ok, n := parser.CastFloatValue(" 5")
	if n != 5 {
		t.Fail()
	}
	ok, _ = parser.CastFloatValue("not a number")
	if ok {
		t.Fail()
	}
}

func TestExtractArray(t *testing.T) {
	ok, scs := parser.ExtractArray("[a,b , c]")
	if !ok || scs[0] != "a" || scs[1] != "b" || scs[2] != "c" {
		t.Error(scs)
	}

}
