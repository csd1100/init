package main

import (
	"fmt"
	"os"

	"github.com/csd1100/init/internal/core"
	"github.com/csd1100/init/internal/utils"
)

func main() {
	args := os.Args[1:]
	options, err := utils.ParseArgs(args)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
	err = core.Init(*options)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
}
