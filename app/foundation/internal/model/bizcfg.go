package model

type Config struct {
	RpcUrl                string  `json:"rpcUrl" v:"required"`
	FaucetAccount         string  `json:"faucetAccount" v:"required"`
	FaucetPassphrase      string  `json:"faucetPassphrase" v:"required"`
	AmountPerTime         float64 `json:"amountPerTime" v:"required"`
	AmountLimitPerAddress float64 `json:"amountLimitPerAddress" v:"required"`
	AmountLimitPerDay     float64 `json:"amountLimitPerDay" v:"required"`
	ClaimsIntervalHours   int64   `json:"claimsIntervalHours" v:"required"`
	RegistryAddress       string  `json:"registryAddress" v:"required"`
	RegistryAccount       string  `json:"registryAccount" v:"required"`
	RegistryPassphrase    string  `json:"registryPassphrase" v:"required"`
}
