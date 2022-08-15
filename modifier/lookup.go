package modifier

import (
	f "github.com/noam-g4/functional"
)

type Path []string

func TracePath(key string, m map[interface{}]interface{}, p Path) Path {
	for k := range m {
		name := k.(string)
		if key == name {
			return f.ConcatSlices(p, Path{name})
		}
	}
	for k := range m {
		name := k.(string)
		if val, ok := m[name].(map[interface{}]interface{}); ok {
			return TracePath(key, val, f.ConcatSlices(p, Path{name}))
		}
	}
	return Path{}
}
