package config

import "github.com/noam-g4/figure/parser"

type Settings struct {
	FilePath   string
	Prefix     string
	Separator  string
	Convention parser.Mode
}
