package replacer

import (
	"os"
)

func GetEnv(name string) (bool, string) {
	value := os.Getenv(name)
	if value == "" {
		return false, value
	}
	return true, value
}
