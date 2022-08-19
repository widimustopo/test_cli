package cmd

import (
	"fmt"
	"test_cli/assets"
	"test_cli/models"

	"github.com/spf13/cobra"
)

func setChainCode(cmd *cobra.Command, args []string) {
	var (
		chaincode models.Chaincode
		counter   int = 1
	)
	assets.Header()
	fmt.Print("how many chaincode : ")
	fmt.Scanln(&counter)
	for i := 0; i < counter; i++ {
		fmt.Print("name chaincode : ")
		fmt.Scanln(&chaincode.Name)
		fmt.Print("channel : ")
		fmt.Scanln(&chaincode.Channel)
		fmt.Print("directory : ")
		fmt.Scanln(&chaincode.Directory)
		fmt.Print("Endorsement : ")
		fmt.Scanln(&chaincode.Endorsement)
		fmt.Print("Init : ")
		fmt.Scanln(&chaincode.Init)
		fmt.Print("Lang : ")
		fmt.Scanln(&chaincode.Lang)
		fmt.Print("Version : ")
		fmt.Scanln(&chaincode.Version)

		globConf.Chaincodes = append(globConf.Chaincodes, chaincode)
	}
	fmt.Scanln()

	fmt.Println("success input chain code")
	if cmd.Name() == "test_cli" {
		assets.Clear()
		choose(cmd, args)
	}
}
