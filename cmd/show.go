package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func showConfig(cmd *cobra.Command, args []string) {

	b, err := json.MarshalIndent(&globConf, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Result Configuration JSON")
	fmt.Println("=========================")
	os.Stdout.Write(b)
}
