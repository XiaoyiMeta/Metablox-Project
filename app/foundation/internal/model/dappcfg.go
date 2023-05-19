package model

import "time"

type DappCfgInput struct {
	ChainId string `json:"chainId"`
}

type DappCfgOutput struct {
	ContractInfos []*DappContractItem `json:"contractInfos"`
	NodeInfos     []*DappNodeItem     `json:"nodeInfos"`
}

type DappNodeItem struct {
	Id        int       `json:"id"        description:""`
	BizType   string    `json:"bizType"   description:""`
	ChainId   int64     `json:"chainId"   description:""`
	RpcUrl    string    `json:"rpcUrl"    description:""`
	WssUrl    string    `json:"wssUrl"    description:""`
	Disabled  int       `json:"disabled"  description:""`
	CreatedAt time.Time `json:"createdAt" description:""`
	UpdatedAt time.Time `json:"updatedAt" description:""`
	DeletedAt time.Time `json:"-" description:""`
	Remark    string    `json:"remark"    description:""`
}

type DappContractItem struct {
	Id              int       `json:"id"              description:""`
	BizType         string    `json:"bizType"         description:""`
	ContractType    string    `json:"contractType"    description:""`
	Description     string    `json:"description"     description:""`
	ChainId         int64     `json:"chainId"         description:""`
	ContractAddress string    `json:"contractAddress" description:""`
	CreatorAddress  string    `json:"creatorAddress"  description:""`
	TxHash          string    `json:"txHash"          description:""`
	BlockTimestamp  int64     `json:"blockTimestamp"  description:""`
	BlockNumber     int64     `json:"blockNumber"     description:""`
	Disabled        int       `json:"disabled"        description:""`
	CreatedAt       time.Time `json:"createdAt"       description:""`
	UpdatedAt       time.Time `json:"updatedAt"       description:""`
	DeletedAt       time.Time `json:"-"               description:""`
}
