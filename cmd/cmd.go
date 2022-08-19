package cmd

import (
	"github.com/spf13/cobra"
)

var (
	base = &cobra.Command{
		Use:   "test_cli",
		Short: "information",
		Long:  "information of application",
		Run:   globalConfig,
	}

	setconfig = &cobra.Command{
		Use:   "set-config",
		Short: "Configuration",
		Long:  "Set your configuration",
		Run:   globalConfig,
	}

	setchaincode = &cobra.Command{
		Use:   "set-chaincode",
		Short: "Set chaincode",
		Long:  "Set your chaincode",
		Run:   setChainCode,
	}

	setchannel = &cobra.Command{
		Use:   "set-channel",
		Short: "Set channel",
		Long:  "Set your channel",
		Run:   setChannel,
	}

	showCurrentConfig = &cobra.Command{
		Use:   "show-config",
		Short: "Show Configuration",
		Long:  "Show your current configuration",
		Run:   showConfig,
	}

	generate = &cobra.Command{
		Use:   "generate",
		Short: "Generate Config",
		Long:  "Generate your configuration",
		Run:   generateConfig,
	}
)

func Execute() {
	cobra.CheckErr(base.Execute())
}

func init() {
	base.AddCommand(setconfig)
	base.AddCommand(setchaincode)
	base.AddCommand(setchannel)
	base.AddCommand(showCurrentConfig)
	base.AddCommand(generate)
}
