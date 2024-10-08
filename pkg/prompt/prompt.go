package prompt

import (
	"errors"
	"os"

	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	ErrorMsg string
	Label    string
}

func PromptGetInput(pc PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    pc.Label,
		Validate: validate,
		Templates: &promptui.PromptTemplates{
			Success: "Display Name: ",
		},
	}

	result, err := prompt.Run()

	if err != nil {
		os.Exit(1)
	}

	return result
}

func PromptGetSelect(pc PromptContent, items []string) string {

	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Items: items,
			Templates: &promptui.SelectTemplates{
				Label:    pc.Label,
				Selected: "Cluster: {{ . }}",
			},
		}

		index, result, err = prompt.Run()

		if err != nil {
			os.Exit(1)
		}

		if index == -1 {
			items = append(items, result)
		}
	}

	return result
}
