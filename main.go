package main

import (
	"github.com/csd1100/init/internal/core"
	"github.com/csd1100/init/internal/helpers"
	"github.com/csd1100/init/internal/utils"
)

func main() {
	helpers.AppLogger.CurrentLevel = helpers.WARN_LEVEL
	helpers.AppLogger.Trace("Starting init....")
	options, err := utils.ParseArgs()
	if err != nil {
		helpers.AppLogger.Error(err.Error())
	}

	if options.Help {
		return
	}

	err = core.Init(*options)
	if err != nil {
		helpers.AppLogger.Error(err.Error())
	}
}
