## noam-g4/figure
A simple, elegant and highly flexible solution for managing your ***go app*** configurations environment.
This package lets you manage your go app config with a ***.yml*** file and overwrite values using ***environment variables*** using a `Replacer` mechanism (more on that later). Essentially, making your app easly configurable when running in different environments.

### Example
Suppose we have this `./config.yml` file:
```yaml
env: test
writeMode: false
retries: 5
```
In our application:
```go
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/noam-g4/figure"
	r "github.com/noam-g4/figure/replacer"
)

type Config struct {
	Env       string `yaml:"env"`
	WriteMode bool   `yaml:"writeMode"`
	Retries   int    `yaml:"retries"`
}

func main() {
	replacers := []r.Replacer[Config]{
		{
			Env: "ENV",
			Setter: func(c Config, s string) Config {
				mod := c
				mod.Env = s
				return mod
			},
		},
		{
			Env: "WRITEMODE",
			Setter: func(c Config, s string) Config {
				mod := c
				if s != "true" {
					mod.WriteMode = false
				}
				mod.WriteMode = true
				return mod
			},
		},
		{
			Env: "RETRIES",
			Setter: func(c Config, s string) Config {
				mod := c
				i, err := strconv.Atoi(s)
				if err != nil {
					return c
				}
				mod.Retries = i
				return mod
			},
		},
	}

	err, config := figure.LoadConfig("./config.yml", replacers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)
}
```
by running our program, our output will be:
```bash
{ test false 5}
```
But, by setting environment variables, our program will replace the specified values from the yaml file by the set environment:
```bash
WRITEMOD=true
RETRIES=7
go run .
{ test true 7 }
```

### Replacer
The `Replacer` data structure, lets you specify the name of the environment variable you want to overwrite with and a special function that takes your config struct and the variable's value so you can choose how to parse this value to your config structure.
This results in a highly flexible mechanism to overwrite config values, given a set of specific environment variables (which works great in cases of *docker compose* and *k8s deployments*)

### Use this package as a simple yaml loader
If you don't need to dynamically change environment variables and you just want a simple package that loads and unmarshall a yaml file, you can use this package as following:
```go
package main

import (
    "fmt"
    "log"

    "github.com/noam-g4/figure"
)

type Config struct {
    Env       string `yaml:"env"`
    WriteMode bool   `yaml:"writeMode"`
    Retries   int    `yaml:"retries"`
}

func main() {
    err, config := figure.LoadConfigWithoutReplacers("./config.yml")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(config)
}
```
