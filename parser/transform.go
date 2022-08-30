package parser

import (
	"strings"

	"github.com/noam-g4/figure/v2/env"
	f "github.com/noam-g4/functional"
)

func toCamel(sep, mSep string, v env.Var) env.Var {
	name := v.Name
	wrds := SmartSplit(name, sep, mSep)
	t := f.Reduce(wrds[1:], func(out, wrd string) string {
		if wrd == "" {
			return out
		}
		w := strings.ToUpper(string(wrd[0])) + strings.ToLower(wrd[1:])
		return out + w
	}, strings.ToLower(wrds[0]))
	return env.Var{
		Name:  t,
		Value: v.Value,
	}
}

func toSnake(sep, mSep string, v env.Var) env.Var {
	name := v.Name
	wrds := SmartSplit(name, sep, mSep)
	t := f.Reduce(wrds[1:], func(out, wrd string) string {
		return out + "_" + strings.ToLower(wrd)
	}, strings.ToLower(wrds[0]))
	return env.Var{
		Name:  t,
		Value: v.Value,
	}
}

func SmartSplit(str, sep, glu string) []string {
	xs := strings.Split(str, sep)
	if len(xs) < 2 {
		return xs
	}
	res := f.Reduce(xs[1:], func(y, x string) string {
		if x == "" {
			return y + glu
		}
		if len(y) >= len(glu) &&
			y[len(y)-len(glu):] == glu {
			return y + x
		}
		return y + " " + x
	}, xs[0])
	return strings.Split(res, " ")
}
