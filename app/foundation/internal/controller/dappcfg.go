package controller

import (
	"context"
	v1 "go-metablox/app/foundation/api/v1"
	"go-metablox/app/foundation/internal/model"
	"go-metablox/app/foundation/internal/service"
)

var Index = &cIndex{}

type cIndex struct {
}

func (c *cIndex) Config(ctx context.Context, req *v1.DappCfgReq) (*v1.DappCfgRes, error) {
	cfg, err := service.DappCfg().GetAllCfg(ctx, model.DappCfgInput{
		ChainId: req.ChainId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DappCfgRes{
		ContractInfos: cfg.ContractInfos,
		NodeInfos:     cfg.NodeInfos,
	}, nil
}
