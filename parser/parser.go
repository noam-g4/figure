package parser

import (
	"strings"

	"github.com/noam-g4/figure/v2/env"
	"gopkg.in/yaml.v2"
)

type Mode int

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
	case 0:
		return toCamel(seperator, v)
	case 1:
		return toSnake(seperator, v)
	case 2:
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

func SerializeYamlMap(m map[interface{}]interface{}) (error, []byte) {
	data, err := yaml.Marshal(&m)
	if err != nil {
		return err, nil
	}
	return nil, data
}
