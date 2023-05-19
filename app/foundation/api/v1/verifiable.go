package v1

import (
	"github.com/MetaBloxIO/did-sdk-go"
	"github.com/gogf/gf/v2/frame/g"
)

type VerifyVPReq struct {
	g.Meta `tags:"Verifiable" summary:"Verify vp"`
	did.VerifiablePresentation
}

type VerifyVPRes struct {
}

type VCNonceReq struct {
	g.Meta `tags:"Verifiable" summary:"Nonce"`
	Did    string `json:"did" v:"required|did"`
}

type VCNonceRes struct {
	Nonce     string `json:"nonce"`
	ExpiresAt string `json:"expiresAt"`
}

type IssueWifiVCReq struct {
	g.Meta    `tags:"Verifiable" summary:"Issue wifi vc"`
	Id        string `json:"id" v:"required|did"`
	Type      string `json:"type" v:"required|in:Validator"`
	Nonce     string `json:"nonce" v:"required"`
	Signature string `json:"signature" v:"required"`
}

type IssueWifiVCRes struct {
	*did.VerifiableCredential
}

type RenewWifiVCReq struct {
	g.Meta `tags:"Verifiable" summary:"Renew wifi vc"`
	did.VerifiablePresentation
}

type RenewWifiVCRes struct {
	*did.VerifiableCredential
}

type RevokeWifiVCReq struct {
	g.Meta `tags:"Verifiable" summary:"Revoke wifi vc"`
	did.VerifiablePresentation
}

type RevokeWifiVCRes struct {
}

type IssueMiningVCReq struct {
	g.Meta    `tags:"Verifiable" summary:"Issue mining vc"`
	Id        string `json:"id" v:"required|did"`
	Type      string `json:"type" v:"required|in:Miner"`
	Nonce     string `json:"nonce" v:"required"`
	Signature string `json:"signature" v:"required"`
}

type IssueMiningVCRes struct {
	*did.VerifiableCredential
}

type RenewMiningVCReq struct {
	g.Meta `tags:"Verifiable" summary:"Renew mining vc"`
	did.VerifiablePresentation
}

type RenewMiningVCRes struct {
	*did.VerifiableCredential
}

type RevokeMiningVCReq struct {
	g.Meta `tags:"Verifiable" summary:"Revoke mining vc"`
	did.VerifiablePresentation
}

type RevokeMiningVCRes struct {
}
