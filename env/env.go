package env

import (
	"os"
	"strings"

	f "github.com/noam-g4/functional"
)

type Var struct {
	Name  string
	Value interface{}
}

func ListEnvs(prefix string) []string {
	envs := os.Environ()
	pfLen := len(prefix)

	return f.Filter(envs, func(e string) bool {
		if len(e) <= pfLen {
			return false
		}
		if e[:pfLen] != prefix {
			return false
		}
		return true
	})
}

func GetEnvsWithValue(envs []string) []Var {
	vars := f.Map(envs, func(e string) Var {
		v := splitKeyVal(e)
		if v.Value != "" {
			return v
		}
		return Var{}
	})

	return f.Filter(vars, func(v Var) bool {
		if v.Value == "" {
			return false
		}
		return true
	})
}

func splitKeyVal(e string) Var {
	strs := strings.Split(e, "=")
	return Var{
		Name:  strs[0],
		Value: strings.Join(strs[1:], "="),
	}
}
