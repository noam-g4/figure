package parser

import (
	"strings"

	"github.com/noam-g4/figure/env"
	f "github.com/noam-g4/functional"
)

func toCamel(sep string, v env.Var) env.Var {
	name := v.Name
	wrds := strings.Split(name, sep)
	t := f.Reduce(wrds[1:], func(out, wrd string) string {
		w := strings.ToUpper(wrd[:1]) + strings.ToLower(wrd[1:])
		return out + w
	}, strings.ToLower(wrds[0]))
	return env.Var{
		Name:  t,
		Value: v.Value,
	}
}

func toSnake(sep string, v env.Var) env.Var {
	name := v.Name
	wrds := strings.Split(name, sep)
	t := f.Reduce(wrds[1:], func(out, wrd string) string {
		return out + "_" + strings.ToLower(wrd)
	}, strings.ToLower(wrds[0]))
	return env.Var{
		Name:  t,
		Value: v.Value,
	}
}
