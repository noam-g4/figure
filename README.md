## noam-g4/figure
[![Go Test](https://github.com/noam-g4/figure/actions/workflows/test.yml/badge.svg)](https://github.com/noam-g4/figure/actions/workflows/test.yml) <br/>
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

	"github.com/noam-g4/figure/v2"
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
	err, conf := figure.LoadConfig[Conf](figure.Settings{
		FilePath:   "./config.yml",
		Prefix:     "MYAPP_",
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
export MYAPP_WRITE_MODE=true
export MYAPP_OTHERS__RETRIES=7
export MYAPP_OTHERS__OPTIONS=[x,y]

go run .
{test true {7 [x y]}}
```
