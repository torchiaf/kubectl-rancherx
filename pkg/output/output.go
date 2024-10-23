package output

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"gopkg.in/yaml.v2"
)

type Table[T any] struct {
	Header []string
	Row    func(item T) []string
}

func PrintTable[T any](ctx context.Context, items []T, table Table[T]) error {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)

	for _, col := range table.Header {
		fmt.Fprintf(w, "%s\t", strings.ToUpper(col))
	}
	fmt.Fprint(w, "\n")

	for _, item := range items {
		for _, col := range table.Row(item) {
			fmt.Fprintf(w, "%s\t", col)
		}
		fmt.Fprint(w, "\n")
	}

	w.Flush()

	return nil
}

func Print[T any](ctx context.Context, outputType string, items []T, table Table[T]) error {
	var content any = items

	if len(items) == 1 {
		content = items[0]
	}

	switch outputType {
	case "yaml":
		yamlData, err := yaml.Marshal(content)
		if err != nil {
			return err
		}

		fmt.Println(string(yamlData))
	case "json":
		jsonData, err := json.MarshalIndent(content, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(jsonData))
	default:
		if len(items) > 0 {
			PrintTable(ctx, items, table)
		}
	}

	return nil
}
