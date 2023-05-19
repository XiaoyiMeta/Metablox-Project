package model

import "github.com/MetaBloxIO/did-sdk-go"

type VCNonceInput struct {
	Did string `json:"did"`
}
type VCChallengeOutput struct {
	ExpiresAt string `json:"expiresAt"`
	Nonce     string `json:"nonce"`
}

type IssueVCInput struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type IssueVCOutput struct {
	*did.VerifiableCredential
}

type VCInput struct {
	did.VerifiablePresentation
}

type VCOutput struct {
	*did.VerifiableCredential
}
