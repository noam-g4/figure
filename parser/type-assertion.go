package parser

import (
	"strconv"
	"strings"

	"github.com/noam-g4/figure/env"
	f "github.com/noam-g4/functional"
)

func CastBooleanValue(val string) (bool, bool) {
	str := strings.ToLower(strings.Trim(val, " "))
	if str == "true" {
		return true, true
	}
	if str == "false" {
		return true, false
	}
	return false, false
}

func CastIntValue(val string) (bool, int) {
	str := strings.Trim(val, " ")
	if n, err := strconv.Atoi(str); err == nil {
		return true, n
	}
	return false, 0
}

func CastFloatValue(val string) (bool, float64) {
	str := strings.Trim(val, " ")
	if n, err := strconv.ParseFloat(str, 64); err == nil {
		return true, n
	}
	return false, 0
}

// enforce array to be in brackets & comma separator
func ExtractArray(val string) (bool, []string) {
	str := strings.Trim(val, " ")
	if len(str) > 0 && str[0] == '[' && str[len(str)-1] == ']' {
		slc := strings.Split(str[1:len(str)-1], ",")
		return true, f.Map(slc, func(s string) string {
			return strings.Trim(s, " ")
		}, f.EmptySet[string]())
	}
	return false, []string{}
}

func InterpretType(value string) interface{} {
	isBool, b := CastBooleanValue(value)
	if isBool {
		return b
	}
	isInt, i := CastIntValue(value)
	if isInt {
		return i
	}
	isFloat, f := CastFloatValue(value)
	if isFloat {
		return f
	}
	return value
}

func InterpretTypeWithArray(value string) interface{} {
	isArray, arr := ExtractArray(value)
	if isArray {
		res := f.Map(arr,
			func(s string) interface{} {
				return InterpretType(s)
			}, f.EmptySet[interface{}]())
		return res
	}
	return InterpretType(value)
}

func AssertVariableValue(v env.Var) env.Var {
	return env.Var{
		Name:  v.Name,
		Value: InterpretTypeWithArray(v.Value.(string)),
	}
}
