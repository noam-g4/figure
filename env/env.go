package env

import (
	"os"
	"strings"

	f "github.com/noam-g4/functional"
)

type Var struct {
	Name  string
	Value string
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
	}, f.EmptySet[string]())
}

func GetEnvsWithValue(envs []string) []Var {
	vars := f.Map(envs, func(e string) Var {
		v := splitKeyVal(e)
		if v.Value != "" {
			return v
		}
		return Var{}
	}, f.EmptySet[Var]())

	return f.Filter(vars, func(v Var) bool {
		if v.Value == "" {
			return false
		}
		return true
	}, f.EmptySet[Var]())
}

func splitKeyVal(e string) Var {
	strs := strings.Split(e, "=")
	return Var{
		Name:  strs[0],
		Value: strs[1],
	}
}