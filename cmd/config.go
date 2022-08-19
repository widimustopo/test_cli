package cmd

import (
	"fmt"
	"test_cli/models"

	"github.com/spf13/cobra"
)

func config(cmd *cobra.Command, args []string) {
	var (
		global models.Global
	)

	fmt.Println("dsadsadsa", cmd.Name())

FV:
	fmt.Print("Febrication Version (1.4.6 or 2.2.4): ")
	fmt.Scanln(&global.FabricVersion)
	if global.FabricVersion != "1.4.6" && global.FabricVersion != "2.2.4" {
		fmt.Println(fmt.Errorf("wrong version must (1.4.6 or 2.2.4)"))
		goto FV
	}

	fmt.Print("TLS (true or false): ")
	fmt.Scanln(&global.TLS)
	fmt.Print("Monitoring Log (debug or prod): ")
	fmt.Scanln(&global.Monitoring.Loglevel)

	globConf.Global = global

	setChannel(cmd, args)

}
