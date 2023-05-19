package model

type FaucetClaimsInput struct {
	Account string `json:"address"`
}

type FaucetClaimsOutput struct {
	Hash string `json:"hash"`
}
