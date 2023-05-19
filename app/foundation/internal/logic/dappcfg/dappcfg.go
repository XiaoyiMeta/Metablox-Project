package dappcfg

import (
	"context"
	"go-metablox/app/foundation/internal/dao"
	"go-metablox/app/foundation/internal/model"
	"go-metablox/app/foundation/internal/service"
	"go-metablox/pkg/gconsts"
)

type sDappCfg struct {
}

func init() {
	service.RegisterDappCfg(New())
}

func New() *sDappCfg {
	return &sDappCfg{}
}

func (s *sDappCfg) GetAllCfg(ctx context.Context, in model.DappCfgInput) (*model.DappCfgOutput, error) {

	contracts := make([]*model.DappContractItem, 0)
	if err := dao.DappContract.Ctx(ctx).
		Where(dao.DappContract.Columns().Disabled, gconsts.False).
		Where(dao.DappContract.Columns().ChainId, in.ChainId).
		OmitEmpty().
		Scan(&contracts); err != nil {
		return nil, err
	}

	nodes := make([]*model.DappNodeItem, 0)
	if err := dao.DappNode.Ctx(ctx).
		Where(dao.DappNode.Columns().Disabled, gconsts.False).
		Where(dao.DappNode.Columns().ChainId, in.ChainId).
		OmitEmpty().
		Scan(&nodes); err != nil {
		return nil, err
	}

	return &model.DappCfgOutput{
		ContractInfos: contracts,
		NodeInfos:     nodes,
	}, nil

}
