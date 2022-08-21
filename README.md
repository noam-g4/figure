## noam-g4/figure
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/noam-g4/figure/Go%20Test?label=tests&style=flat-square) <br/>
A very simple package for loading your Go app with a `.yml` configuration and overwrite its values with environment variables. 

### Example
Suppose we have this `./config.yml` file:
```yaml
env: test
writeMode: false
others:
  retries: 5
  options: 
    - a
    - b
    - c
```
In our application:
```go
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/noam-g4/figure"
	"github.com/noam-g4/figure/parser"
	"github.com/noam-g4/figure/config"
)

type Conf struct {
	Env       string `yaml:"env"`
	WriteMode bool   `yaml:"writeMode"`
	Others    struct {
		Retries int      `yaml:"retries"`
		Options []string `yaml:"options"`
	} `yaml:"others"`
}

func main() {
	err, conf := figure.LoadConfig[Conf](config.Settings{
		FilePath:   "./config.yml",
		Prefix:     "MYAPP_",
		Convention: parser.Camel,
		Separator:  "_",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf)
}
```
by running our program, our output will be:
```bash
{test false {5 [a b c]}}
```
But, by setting environment variables, our program will replace the specified values from the yaml file by the set environment:
```bash
export MAYAPP_WRITE_MOD=true
export MYAPP_RETRIES=7
export MYAPP_OPTIONS=[x,y]

go run .
{test true {7 [x y]}}
```
