package fetcher

import (
	"io/ioutil"
)

func ReadFile(path string) (error, []byte) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}
	return nil, data
}
