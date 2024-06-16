package flag

import (
	"errors"
	"fmt"
	"strings"
)

type SetFlag string

func (e *SetFlag) String() string {
	return string(*e)
}

func (e *SetFlag) Set(v string) error {

	output := map[string]string{}

	kv := strings.Split(v, "=")

	if len(kv) != 2 {
		return errors.New(`must be "key=value" format`)
	}

	output[kv[0]] = kv[1]

	fmt.Println(output)

	*e = SetFlag(v)

	return nil
}

func (e *SetFlag) Type() string {
	return "key=value"
}
