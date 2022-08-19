package cmd

import (
	"fmt"
	"test_cli/models"

	"github.com/spf13/cobra"
)

func setChannel(cmd *cobra.Command, args []string) {
	var (
		peerstemp string = ""
		channel   models.Channel
		org       models.Org
		counterC  int = 1
		counterP  int = 1
		counterO  int = 1
	)

	fmt.Println("how many channels : ")
	fmt.Scan(&counterC)
	for i := 0; i < counterC; i++ {
		fmt.Println("channel name : ")
		fmt.Scan(&channel.Name)
		fmt.Println("how many organizations : ")
		fmt.Scan(&counterO)
		for i := 0; i < counterO; i++ {
			fmt.Println("org name : ")
			fmt.Scan(&org.Name)
			fmt.Println("how many peers : ")
			fmt.Scan(&counterP)
			for i := 0; i < counterP; i++ {
				fmt.Println("peers name : ")
				fmt.Scan(&peerstemp)
				org.Peers = append(org.Peers, peerstemp)
			}
			channel.Orgs = append(channel.Orgs, org)
		}
		globConf.Channels = append(globConf.Channels, channel)
	}

	setChainCode(cmd, args)
}
