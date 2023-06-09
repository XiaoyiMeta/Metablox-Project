// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"go-metablox/app/foundation/internal/model"
)

type (
	IDappCfg interface {
		GetAllCfg(ctx context.Context, in model.DappCfgInput) (*model.DappCfgOutput, error)
	}
)

var (
	localDappCfg IDappCfg
)

func DappCfg() IDappCfg {
	if localDappCfg == nil {
		panic("implement not found for interface IDappCfg, forgot register?")
	}
	return localDappCfg
}

func RegisterDappCfg(i IDappCfg) {
	localDappCfg = i
}
