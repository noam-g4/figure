package parser

import (
	"strings"

	"github.com/noam-g4/figure/env"
	"gopkg.in/yaml.v2"
)

type Mode int

const (
	Camel     Mode = 0
	Snake     Mode = 1
	Caps      Mode = 2
	SnakeCaps Mode = 3
)

func StripPrefix(prefix string, v env.Var) env.Var {
	return env.Var{
		Name:  v.Name[len(prefix):],
		Value: v.Value,
	}
}

func TransformName(mode Mode, seperator string, v env.Var) env.Var {
	if mode > 3 {
		return v
	}
	switch mode {
	case Camel:
		return toCamel(seperator, v)
	case Snake:
		return toSnake(seperator, v)
	case Caps:
		return env.Var{
			Name:  strings.ToUpper(v.Name),
			Value: v.Value,
		}
	default:
		y := toSnake(seperator, v)
		return env.Var{
			Name:  strings.ToUpper(y.Name),
			Value: y.Value,
		}
	}
}

func ParseToMap(b []byte) (error, map[interface{}]interface{}) {
	var m map[interface{}]interface{}
	err := yaml.Unmarshal(b, &m)
	if err != nil {
		return err, m
	}
	return nil, m
}

func Parse[T interface{}](b []byte) (error, T) {
	var c T
	err := yaml.Unmarshal(b, &c)
	if err != nil {
		return err, c
	}
	return nil, c
}
