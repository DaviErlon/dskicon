package menu

import (
	"github.com/DaviErlon/dskicon/util"
	"strings"

	"github.com/manifoldco/promptui"
)

func MdfDesktop(app *util.AppConfig) (State, error) {

	prompt := promptui.Prompt{
		Label:       "",
		Validate:    Valid,
		HideEntered: true,
		Default:     app.DesktopDir,
	}

	path, err := prompt.Run()
	if err != nil {
		return BASE, err
	}

	path = strings.TrimSpace(RemoveNonPrintable(path))

	if path != "" && path != app.DesktopDir {
		err = app.SetDesktopDir(path)
	}

	return CONFIG, err
}
