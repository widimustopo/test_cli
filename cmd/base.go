package cmd

import (
	"test_cli/assets"
	"test_cli/models"

	"github.com/spf13/cobra"
)

var globConf models.Configurations

func baseFunc(cmd *cobra.Command, args []string) {
	assets.Header()
	choose(cmd, args)
}
