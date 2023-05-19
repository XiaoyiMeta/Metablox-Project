package v1

import (
	"github.com/MetaBloxIO/did-sdk-go"
	"github.com/gogf/gf/v2/frame/g"
)

type WifiUserInfoReq struct {
	g.Meta ` tags:"Wifi" summary:"User Info"`
	did.VerifiablePresentation
}

type WifiUserInfoRes struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WifiCertFileReq struct {
	g.Meta ` tags:"Wifi" summary:"CertFile"`
}

type WifiCertFileRes struct {
}
