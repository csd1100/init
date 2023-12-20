package main

import (
	"github.com/csd1100/init/internal/core"
	"github.com/csd1100/init/internal/helpers"
	"github.com/csd1100/init/internal/utils"
)

func main() {
	helpers.AppLogger.Trace("Starting init....")
	options, err := utils.ParseArgs()
	if options.Help {
		return
	}
	if err != nil {
		helpers.AppLogger.Panic(err.Error())
	}

	err = core.Init(*options)
	if err != nil {
		helpers.AppLogger.Error(err.Error())
	}
}
