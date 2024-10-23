package output

import (
	"context"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

func Print[T any](ctx context.Context, item T, formatter string, def func(item T) string) error {
	switch formatter {
	case "yaml":
		yamlData, err := yaml.Marshal(item)
		if err != nil {
			return err
		}

		fmt.Println(string(yamlData))
	case "json":
		jsonData, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(jsonData))
	default:
		fmt.Println(def(item))
	}

	return nil
}
