package models

import "encoding/json"

type Global struct {
	FabricVersion string     `json:"fabricVersion"`
	TLS           bool       `json:"tls"`
	Monitoring    Monitoring `json:"monitoring"`
}

type Monitoring struct {
	Loglevel string `json:"loglevel"`
}

type Channel struct {
	Name string `json:"name"`
	Orgs []Org  `json:"orgs"`
}

type Org struct {
	Name  string   `json:"name"`
	Peers []string `json:"peers"`
}

type Chaincode struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Lang        string `json:"lang"`
	Channel     string `json:"channel"`
	Init        string `json:"init"`
	Endorsement string `json:"endorsement"`
	Directory   string `json:"directory"`
}

type Configurations struct {
	Schema     string           `json:"$schema"`
	Global     Global           `json:"global"`
	Orgs       *json.RawMessage `json:"orgs"`
	Channels   []Channel        `json:"channels"`
	Chaincodes []Chaincode      `json:"chaincodes"`
}
