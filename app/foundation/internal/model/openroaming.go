package model

type CreateWifiUserInfoInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WifiUserInfoOutput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
