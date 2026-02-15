package menu

import "github.com/manifoldco/promptui"

func Config() (State, error) {
	options := []Options{
		{"\U00002190 Back", "Return to the main menu"},
		{"Desktop Directory", "Change the desktop folder where icons are created"},
		{"Search Directories", "Manage directories used to search for application icons"},
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
		result = BASE
	case 1:
		result = MDF_DESKTOP
	case 2:
		result = SEARCHDIR
	}

	return result, err
}