package menu

import (
	"github.com/DaviErlon/dskicon/util"
	"strings"

	"github.com/manifoldco/promptui"
)

func Icons(app *util.AppConfig, result int) (State, error) {
	icons, err := app.GetIcons(result)
	if err != nil {
		return 0, err
	}
	defer icons.Close()

	itens := append([]util.IconFile{
		{Name: "\U00002190 Back", OriginName: "Return to the Add menu"}},
		icons...,
	)

	searcher := func(input string, index int) bool {
		if index == 0 {
			return false
		}

		icon := icons[index-1]

		name := strings.ReplaceAll(strings.ToLower(icon.Name), " ", "")
		input = strings.ReplaceAll(strings.ToLower(input), " ", "")

		return strings.Contains(name, input)
	}

	templates := &promptui.SelectTemplates{
		Label:    " {{ . | white | bold }}",
		Active:   "\U0001F9E9 {{ .Name | green }}",
		Inactive: "   {{ .Name |  white }}",
		Details: `
{{ .OriginName | faint }}`,
	}

	prompt := promptui.Select{
		Label:        "Select an icon:",
		Items:        itens,
		Templates:    templates,
		Searcher:     searcher,
		Size:         10,
		HideSelected: true,
	}

	result, _, err = prompt.Run()
	if result == 0 {
		return ADD, err
	}

	err = icons.AddIconsToDesktop(result-1, app.DesktopDir)

	return ICONS, err
}
