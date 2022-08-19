package cmd

import (
	"fmt"
	"os"
	"test_cli/assets"

	"github.com/spf13/cobra"
)

func choose(cmd *cobra.Command, args []string) {
	var (
		option string
	)
MENU:
	assets.Menu()
	fmt.Print("input : ")
	fmt.Scanln(&option)
	switch option {
	case "0":
		os.Exit(0)
	case "1":
		assets.Clear()
		baseFunc(cmd, args)
	case "2":
		assets.Clear()
		config(cmd, args)
	case "3":
		assets.Clear()
		setChainCode(cmd, args)
	case "4":
		assets.Clear()
		setChannel(cmd, args)
	case "5":
		assets.Clear()
		showConfig(cmd, args)
	case "6":
		assets.Clear()
		generateConfig(cmd, args)
	default:
		assets.Clear()
		goto MENU
	}

}
