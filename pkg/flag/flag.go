package flag

import (
	"errors"
	"strings"
)

type Set map[string]string

func (e *Set) String() string {
	return string("")
}

func (e *Set) Set(v string) error {

	output := make(map[string]string)

	kv := strings.Split(v, "=")

	if len(kv) != 2 {
		return errors.New(`must be "key=value" format`)
	}

	if *e != nil {
		for k, v := range *e {
			output[k] = v
		}
	}

	output[kv[0]] = kv[1]

	*e = Set(output)

	return nil
}

func (e *Set) Type() string {
	return "key=value"
}
