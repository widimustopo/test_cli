package cmd

import (
	"fmt"
	"test_cli/models"

	"github.com/spf13/cobra"
)

func setChainCode(cmd *cobra.Command, args []string) {
	var (
		chaincode models.Chaincode
		counter   int = 1
	)

	fmt.Println("how many chaincode : ")
	fmt.Scan(&counter)
	for i := 0; i < counter; i++ {
		fmt.Println("name chaincode : ")
		fmt.Scan(&chaincode.Name)
		fmt.Println("channel : ")
		fmt.Scan(&chaincode.Channel)
		fmt.Println("directory : ")
		fmt.Scan(&chaincode.Directory)
		fmt.Println("Endorsement : ")
		fmt.Scan(&chaincode.Endorsement)
		fmt.Println("Init : ")
		fmt.Scan(&chaincode.Init)
		fmt.Println("Lang : ")
		fmt.Scan(&chaincode.Lang)
		fmt.Println("Version : ")
		fmt.Scan(&chaincode.Version)

		globConf.Chaincodes = append(globConf.Chaincodes, chaincode)
	}

	showConfig(cmd, args)
}
