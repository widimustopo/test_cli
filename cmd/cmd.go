package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

var (
	base = &cobra.Command{
		Use:   "test_cli",
		Short: "information",
		Long:  "information of application",
		Run:   baseFunc,
	}

	setconfig = &cobra.Command{
		Use:   "set-config",
		Short: "Configuration",
		Long:  "Set your configuration",
		Run:   config,
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
	globConf.Schema = "https://github.com/hyperledger/releases/download/1.1.0/schema.json"
	orgs := json.RawMessage(`[ { "organization": { "name": "Orderer", "domain": "orderer.example.com" }, "orderers": [ { "groupName": "group1", "prefix": "orderer", "type": "raft", "instances": 1 } ] }, { "organization": { "name": "Org1", "mspName": "Org1MSP", "domain": "org1.example.com" }, "ca": { "prefix": "ca" }, "peer": { "prefix": "peer", "instances": 2, "db": "LevelDb" } } ]`)
	globConf.Orgs = &orgs

	cobra.CheckErr(base.Execute())

}

func init() {
	base.AddCommand(setconfig)
	base.AddCommand(setchaincode)
	base.AddCommand(setchannel)
	base.AddCommand(showCurrentConfig)
	base.AddCommand(generate)
}
