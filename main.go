package main

import (
	"fmt"

	"github.com/csd1100/init/internal/core"
	"github.com/csd1100/init/internal/utils"
)

func main() {
	options, err := utils.ParseArgs()
	if err != nil {
		panic(fmt.Errorf("error: %s\n", err.Error()))
	}

	if options.Help {
		return
	}

	err = core.Init(*options)
	if err != nil {
		panic(fmt.Errorf("error: %s\n", err.Error()))
	}
}
