package figure

import (
	"github.com/noam-g4/figure/config"
	"github.com/noam-g4/figure/env"
	"github.com/noam-g4/figure/fetcher"
	"github.com/noam-g4/figure/modifier"
	"github.com/noam-g4/figure/parser"
	f "github.com/noam-g4/functional"
)

func LoadConfig[C interface{}](s config.Settings) (error, C) {
	var c C
	err, file := fetcher.ReadFile(s.FilePath)
	if err != nil {
		return err, c
	}

	err, yamlMap := parser.ParseToMap(file)
	if err != nil {
		return err, c
	}

	envs := env.GetEnvsWithValue(env.ListEnvs(s.Prefix))

	envsNoPf := f.Map(envs, func(v env.Var) env.Var {
		return parser.StripPrefix(s.Prefix, v)
	}, f.EmptySet[env.Var]())

	transformedEnvs := f.Map(envsNoPf, func(v env.Var) env.Var {
		return parser.TransformName(s.Convention, s.Separator, v)
	}, f.EmptySet[env.Var]())

	updatedMap := modifier.UpdateMapWithEnvs(transformedEnvs, yamlMap)

	err, srlz := parser.SerializeYamlMap(updatedMap)
	if err != nil {
		return err, c
	}

	err, out := parser.Parse[C](srlz)
	if err != nil {
		return err, c
	}
	return nil, out
}
