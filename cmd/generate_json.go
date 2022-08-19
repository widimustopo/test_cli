package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"test_cli/assets"

	"github.com/spf13/cobra"
)

func generateConfig(cmd *cobra.Command, args []string) {
	file, _ := json.MarshalIndent(&globConf, "", " ")

	err := ioutil.WriteFile("./sample.json", file, 0644)
	if err != nil {
		fmt.Println(fmt.Errorf("something when wrong, when created a json file"))
	}
	fmt.Println("Success Generate sample.json")
	fmt.Scanln()
	if cmd.Name() == "test_cli" {
		assets.Clear()
		choose(cmd, args)
	}
}
