package menu

import (
	"github.com/DaviErlon/dskicon/util"
	"strings"

	"github.com/manifoldco/promptui"
)

func MdfSearchDir(app *util.AppConfig, result int) (State, error) {

	if result == -1 {
		name, path, err := mdfaux()
		if err != nil {
			return BASE, err
		}
		err = app.AddSearchDir(name, path)
		if err != nil {
			return BASE, err
		}
		return CONFIG, nil
	}

	namedefault := app.SearchDirs[result].Name
	pathdefault := app.SearchDirs[result].Path

	prompt := promptui.Prompt{
		Label:       namedefault,
		HideEntered: true,
		Default:     pathdefault,
	}

	path, err := prompt.Run()
	if err != nil {
		return BASE, err
	}

	path = strings.TrimSpace(RemoveNonPrintable(path))

	if path != pathdefault {
		if path != "" {
			err = app.UpdateSearchDir(result, path)
		} else {
			err = app.RemoveSearchDir(result)
		}
	}

	return CONFIG, err
}

func mdfaux() (string, string, error) {
	prompt := promptui.Prompt{
		Label:       "Name",
		Validate:    Valid2,
		HideEntered: true,
	}

	name, err := prompt.Run()
	if err != nil {
		return "", "", err
	}

	prompt.Label = "Path for " + name
	prompt.Validate = Valid

	path, err := prompt.Run()

	return name, path, err
}
