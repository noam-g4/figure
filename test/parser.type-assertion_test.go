package test

import (
	"testing"

	"github.com/noam-g4/figure/v2/env"
	"github.com/noam-g4/figure/v2/parser"
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

func TestInterpretType(t *testing.T) {
	b := parser.InterpretType("FALSE")
	i := parser.InterpretType("50")
	f := parser.InterpretType("3.14")
	s := parser.InterpretType("string")

	if _, ok := b.(bool); !ok {
		t.Error(b)
	}
	if _, ok := i.(int); !ok {
		t.Error(i)
	}
	if _, ok := f.(float64); !ok {
		t.Error(f)
	}
	if _, ok := s.(string); !ok {
		t.Error(s)
	}
}

func TestInterpretTypeWithArray(t *testing.T) {
	i := parser.InterpretTypeWithArray("1")
	if i != 1 {
		t.Fail()
	}
	arrInt := parser.InterpretTypeWithArray("[1, 2, 3]")
	aInt := arrInt.([]interface{})
	if len(aInt) != 3 && aInt[2] != 3 {
		t.Error(arrInt)
	}
	compArr := parser.InterpretTypeWithArray("[a, 1, true]")
	cArr := compArr.([]interface{})
	if cArr[0] != "a" && cArr[1] != 1 && !cArr[2].(bool) {
		t.Error(cArr)
	}
}

func TestAssertVariableValue(t *testing.T) {
	v := parser.AssertVariableValue(env.Var{
		Name:  "RETRIES",
		Value: "5",
	})
	if v.Value != 5 {
		t.Fail()
	}
}

func TestAssertURIToString(t *testing.T) {
	uri := "mongodb://admin:pass@docdb-cluster.us-east-1.docdb.amazonaws.com:27017/?retryWrites=false&ssl=true&sslCertificateAuthorityFile=secret/cert_ddb.pem&maxIdleTimeMS=30000"
	v := parser.AssertVariableValue(env.Var{
		Name:  "URI",
		Value: uri,
	})

	if v.Value != uri {
		t.Error(v.Value)
	}
}
