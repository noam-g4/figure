package figure

import (
	"github.com/noam-g4/figure/fetcher"
	"github.com/noam-g4/figure/parser"
	"github.com/noam-g4/figure/replacer"
)

func LoadConfig[C interface{}](
	filepath string,
	rs []replacer.Replacer[C],
) (error, C) {

	err, conf := LoadConfigWithoutReplacers[C](filepath)
	if err != nil {
		return err, conf
	}

	return nil, replacer.ReplaceConfigWithEnv(conf, rs)

}

func LoadConfigWithoutReplacers[C interface{}](
	filePath string,
) (error, C) {

	var e C

	err, data := fetcher.ReadFile(filePath)
	if err != nil {
		return err, e
	}

	err, conf := parser.Parse[C](data)
	if err != nil {
		return err, e
	}

	return nil, conf

}
