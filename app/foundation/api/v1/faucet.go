package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type FaucetClaimsReq struct {
	g.Meta  `tags:"Faucet" summary:"Claims"`
	Account string `v:"required" dc:"account" json:"account"`
}

type FaucetClaimsRes struct {
	TxHash string `json:"txHash" dc:"transaction hash"`
}

type FaucetChallengeReq struct {
	g.Meta `tags:"Faucet" summary:"Challenge"`
	DID    string `v:"required" dc:"did" json:"did"`
}

type FaucetChallengeRes struct {
	UUID string `json:"nonce" dc:"transaction hash"`
}
