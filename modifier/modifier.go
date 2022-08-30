package modifier

import (
	"strings"

	"github.com/noam-g4/figure/v2/env"
	"github.com/noam-g4/figure/v2/parser"
)

func Modify(
	v env.Var,
	m map[interface{}]interface{},
	mSep string,
) map[interface{}]interface{} {
	path := strings.Split(v.Name, mSep)
	mod := m
	for k := range mod {
		val, isMap := mod[k].(map[interface{}]interface{})
		if k.(string) == path[0] && !isMap {
			mod[k] = v.Value
			return mod
		}
		if isMap && path[0] == k.(string) {
			res := Modify(env.Var{
				Name:  strings.Join(path[1:], mSep),
				Value: v.Value,
			}, val, mSep)
			if res != nil {
				mod[k] = res
				return mod
			}
		}
	}
	return mod
}

func ValueOrDefault(sep, def string) string {
	if sep != "" {
		return sep
	}
	return def
}

func UpdateMapWithEnvs(
	envs []env.Var,
	m map[interface{}]interface{},
	mSep string,
) map[interface{}]interface{} {
	if len(envs) == 0 {
		return m
	}

	modEnv := parser.AssertVariableValue(envs[0])
	modMap := Modify(modEnv, m, mSep)
	return UpdateMapWithEnvs(envs[1:], modMap, mSep)
}
