package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"go-metablox/app/foundation/internal/model"
)

type DappCfgReq struct {
	g.Meta  `tags:"Dapp" summary:"Config"`
	ChainId string `json:"chainId"`
}

type DappCfgRes struct {
	g.Meta        `tags:"Dapp Config" summary:"Info"`
	ContractInfos []*model.DappContractItem `json:"contractInfos"`
	NodeInfos     []*model.DappNodeItem     `json:"nodeInfos"`
}
