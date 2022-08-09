package parser

import (
	"gopkg.in/yaml.v2"
)

type any interface{}

func Parse[T any](b []byte) (error, T) {
	var c T
	err := yaml.Unmarshal(b, &c)
	if err != nil {
		return err, c
	}
	return nil, c
}
