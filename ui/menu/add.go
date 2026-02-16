package menu

import (
	"github.com/DaviErlon/dskicon/util"

	"github.com/manifoldco/promptui"
)

func Add(app *util.AppConfig) (State, int, error) {
	templates := &promptui.SelectTemplates{
		Label:    " {{ . | white | bold }}",
		Active:   "\U0001F4C2 {{ .Name | green }}",
		Inactive: "   {{ .Name |  white }}",
		Details: `
{{ .Path | faint }}`,
	}

	itens := append([]util.SearchDirEntry{
		{Name: "\U00002190 Back", Path: "Return to the main menu"}},
		app.SearchDirs...,
	)

	prompt := promptui.Select{
		Label:        "Select an option:",
		Items:        itens,
		Templates:    templates,
		Size:         10,
		HideSelected: true,
	}

	result, _, err := prompt.Run()

	if result == 0 {
		return BASE, 0, err
	}

	return ICONS, result - 1, err
}
