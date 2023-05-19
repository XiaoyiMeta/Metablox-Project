package controller

import (
	"context"
	"fmt"
	"github.com/MetaBloxIO/did-sdk-go"
	"github.com/gogf/gf/v2/frame/g"
	v1 "go-metablox/app/foundation/api/v1"
	"go-metablox/app/foundation/internal/packed"
	"go-metablox/app/foundation/internal/service"
	"net/url"
)

var Wifi = &cWifi{}

type cWifi struct {
}

func (c *cWifi) UserInfo(ctx context.Context, req *v1.WifiUserInfoReq) (*v1.WifiUserInfoRes, error) {

	err := service.Verifiable().VerifyVP(ctx, req.VerifiablePresentation)
	if err != nil {
		return nil, err
	}

	for _, vc := range req.VerifiablePresentation.VerifiableCredential {
		err = service.Verifiable().VerifyVC(ctx, vc)
		if err != nil {
			return nil, err
		}

	}

	dids, _ := did.PrepareDID(req.Holder)
	suffix := dids[2]
	info, err := service.Openroaming().GetUserInfo(ctx, suffix)
	if err != nil {
		return nil, err
	}

	return &v1.WifiUserInfoRes{
		Username: info.Username,
		Password: info.Password,
	}, nil

}

func (c *cWifi) CertFile(ctx context.Context, req *v1.WifiCertFileReq) (*v1.WifiCertFileRes, error) {
	r := g.RequestFromCtx(ctx)
	file, _ := packed.GetWifiServerFile()
	r.Response.Write(file)
	r.Response.Header().Set("Content-Type", "application/force-download")
	r.Response.Header().Set("Accept-Ranges", "bytes")
	r.Response.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename=%s`, url.QueryEscape(packed.WifiServerFileName)))
	return nil, nil
}
