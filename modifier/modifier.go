package modifier

import "github.com/noam-g4/figure/env"

func Modify(v env.Var, m map[interface{}]interface{}) map[interface{}]interface{} {
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
