package ui

import (
	"dskicon/ui/menu"
	"dskicon/util"
)

func Run(app *util.AppConfig) error {
	defer menu.ClearScreen()

	if !app.IsConfigured() {
		dir, err := menu.Welcome()
		if err != nil {
			return err
		}
		err = app.SetDesktopDir(dir)
		if err != nil {
			return err
		}
	}

	menu.ShowBanner()

	var (
		state  menu.State
		err    error
		result int
	)

	for ; ; menu.ShowBanner() {

		if err != nil {
			return err
		}

		switch state {
		case menu.BASE:
			state, err = menu.Base()

		case menu.ADD:
			state, result, err = menu.Add(app)

		case menu.ICONS:
			state, err = menu.Icons(app, result)

		case menu.CONFIG:
			state, err = menu.Config()

		case menu.MDF_DESKTOP:
			state, err = menu.MdfDesktop(app)

		case menu.SEARCHDIR:
			state, result, err = menu.SearchDir(app)

		case menu.MDF_SEARCHDIR:
			state, err = menu.MdfSearchDir(app, result)

		case menu.EXIT:
			fallthrough
		default:
			return nil
		}
	}
}
