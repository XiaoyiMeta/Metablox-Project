package controller

import (
	"context"
	"go-metablox/app/foundation/internal/model"
	"go-metablox/app/foundation/internal/service"

	"go-metablox/app/foundation/api/v1"
)

var (
	Faucet = cFaucet{}
)

type cFaucet struct{}

func (c *cFaucet) Claims(ctx context.Context, req *v1.FaucetClaimsReq) (*v1.FaucetClaimsRes, error) {
	out, err := service.Faucet().Claims(ctx, model.FaucetClaimsInput{
		Account: req.Account,
	})
	if err != nil {
		return nil, err
	}

	return &v1.FaucetClaimsRes{TxHash: out.Hash}, nil
}
