package modifier

import (
	"github.com/noam-g4/figure/env"
)

func FindValue(key string, m map[interface{}]interface{}) *interface{} {
	for k, v := range m {
		if k.(string) == key {
			return &v
		}
		if val, ok := m[k].(map[interface{}]interface{}); ok {
			res := FindValue(key, val)
			if res != nil {
				return res
			}
		}
	}
	return nil
}

func GetModifier(v env.Var, m map[interface{}]interface{}) Modifier {
	return Modifier{
		Var:      v,
		Accessor: FindValue(v.Name, m),
	}
}
