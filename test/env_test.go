package test

import (
	"os"
	"testing"

	"github.com/noam-g4/figure/env"
)

func TestListEnv(t *testing.T) {
	os.Setenv("TST_ONE", "one")
	os.Setenv("TST_TWO", "two")

	envs := env.ListEnvs("TST_")
	if len(envs) < 2 {
		t.Fail()
	}

	empty := env.ListEnvs("NOTEXISTING")
	if len(empty) > 0 {
		t.Fail()
	}
}

func TestGetEnvsWithValue(t *testing.T) {
	envs := env.ListEnvs("TST_")
	invalid := env.ListEnvs("NOTEXISTING")

	vars := env.GetEnvsWithValue(envs)
	empty := env.GetEnvsWithValue(invalid)

	checks := []bool{}
	for _, v := range vars {
		if v.Name == "TST_ONE" && v.Value == "one" {
			checks = append(checks, true)
		}
		if v.Name == "TST_TWO" && v.Value == "two" {
			checks = append(checks, true)
		}
	}

	if len(checks) != 2 {
		t.Fail()
	}

	if len(empty) > 0 {
		t.Fail()
	}
}
