package modifier

import (
	"github.com/noam-g4/figure/v2/env"
	"github.com/noam-g4/figure/v2/parser"
)

func Modify(
	v env.Var,
	m map[interface{}]interface{},
) map[interface{}]interface{} {
	mod := m
	for k := range mod {
		if k.(string) == v.Name {
			mod[k] = v.Value
			return mod
		}
		if val, ok := mod[k].(map[interface{}]interface{}); ok {
			res := Modify(v, val)
			if res != nil {
				mod[k] = res
				return mod
			}
		}
	}
	return nil
}

func UpdateMapWithEnvs(
	envs []env.Var,
	m map[interface{}]interface{},
) map[interface{}]interface{} {
	if len(envs) == 0 {
		return m
	}

	modEnv := parser.AssertVariableValue(envs[0])
	modMap := Modify(modEnv, m)
	return UpdateMapWithEnvs(envs[1:], modMap)
}
