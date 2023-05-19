package cmd

import (
	"go-metablox/app/foundation/internal/consts"
	"go-metablox/app/foundation/internal/controller"
	"go-metablox/app/foundation/internal/service"
	"go-metablox/pkg/biz"
	"go-metablox/pkg/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gsession"
)

func router() *ghttp.Server {
	s := g.Server()
	s.SetSessionStorage(gsession.NewStorageMemory())
	WrapperOpenAPI(s)

	s.Use(middleware.Log, service.Middleware().Ctx, middleware.ResponseHandler)
	v1 := s.Group(consts.AppName)
	v1.POST("/faucet/claims", controller.Faucet.Claims)

	v1.POST("/wifi/userInfo", controller.Wifi.UserInfo)
	v1.GET("/wifi/certFile", controller.Wifi.CertFile)

	v1.POST("/vp/verify", controller.Verifiable.VerifyVP)
	v1.GET("/nonce", controller.Verifiable.Nonce)
	v1.POST("/vc/wifi/issue", controller.Verifiable.IssueWifiVC)
	v1.POST("/vc/wifi/renew", controller.Verifiable.RenewWifiVC)
	v1.POST("/vc/wifi/revoke", controller.Verifiable.RevokeWifiVC)

	v1.POST("/vc/mining/issue", controller.Verifiable.IssueMiningVC)
	v1.POST("/vc/mining/renew", controller.Verifiable.RenewMiningVC)
	v1.POST("/vc/mining/revoke", controller.Verifiable.RevokeMiningVC)

	v1.POST("/config", controller.Index.Config)

	return s
}

func WrapperOpenAPI(s *ghttp.Server) {
	api := s.GetOpenApi()
	api.Config.CommonResponse = &biz.Result{}
	api.Config.CommonResponseDataField = "Data"
}
