package cmd

import (
	"fmt"
	"test_cli/assets"
	"test_cli/models"

	"github.com/spf13/cobra"
)

func config(cmd *cobra.Command, args []string) {
	var (
		global models.Global
	)

	assets.Header()

FV:
	fmt.Print("Febrication Version (1.4.6 or 2.2.4): ")
	fmt.Scanln(&global.FabricVersion)
	if global.FabricVersion != "1.4.6" && global.FabricVersion != "2.2.4" {
		fmt.Println(fmt.Errorf("wrong version must (1.4.6 or 2.2.4)"))
		goto FV
	}

	fmt.Print("TLS (true or false): ")
	fmt.Scanln(&global.TLS)

LOG:
	fmt.Print("Monitoring Log (enable or disable): ")
	fmt.Scanln(&global.Monitoring.Loglevel)
	if global.Monitoring.Loglevel != "enable" && global.Monitoring.Loglevel != "disable" {
		fmt.Println(fmt.Errorf("wrong input must (enable or disable)"))
		goto LOG
	}

	globConf.Global = global

	fmt.Scanln()

	fmt.Println("success input config :", global)
	if cmd.Name() == "test_cli" {
		assets.Clear()
		choose(cmd, args)
	}

}
