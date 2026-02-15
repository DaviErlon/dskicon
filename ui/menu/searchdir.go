package menu

import (
	"dskicon/util"

	"github.com/manifoldco/promptui"
)

func SearchDir(app *util.AppConfig) (State, int, error) {

	templates := &promptui.SelectTemplates{
		Label:    " {{ . | white | bold }}",
		Active:   "\U0001F4C2 {{ .Name | green }}",
		Inactive: "   {{ .Name |  white }}",
		Details: `
{{ .Path | faint }}`,
	}

	itens := append([]util.SearchDirEntry{
		{Name: "\U00002190 Back", Path: "Return to the Config menu"},
		{Name: "+ New Path", Path: "Add new directory for application icons"}},
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
		return CONFIG, 0, err
	}

	return MDF_SEARCHDIR, result - 2, err
}
