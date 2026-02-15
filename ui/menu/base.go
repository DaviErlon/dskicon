package menu

import "github.com/manifoldco/promptui"

func Base() (State, error) {
	options := []Options{
		{"Add", "Add an application icon to the desktop"},
		{"Config", "Configure desktop and application icon directories"},
		{"Exit", "Exit dskicon"},
	}

	templates := &promptui.SelectTemplates{
		Label:    " {{ . | white | bold }}",
		Active:   "\U0001F4BB {{ .Name | green }}",
		Inactive: "   {{ .Name |  white }}",
		Details: `
{{ .Details | faint }}`,
	}

	prompt := promptui.Select{
		Label:        "Select an option:",
		Items:        options,
		Templates:    templates,
		HideSelected: true,
	}

	result, _, err := prompt.Run()

	switch result {
	case 0:
		result = ADD
	case 1:
		result = CONFIG
	case 2:
		result = EXIT
	}

	return result, err
}