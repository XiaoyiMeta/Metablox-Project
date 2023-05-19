package validator

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/util/gvalid"
	"go-metablox/pkg/util/ethutil"
	"go-metablox/pkg/util/regutil"
)

func init() {
	gvalid.RegisterRule("eth", RuleETHAddress)
	gvalid.RegisterRule("tx", RuleTxHash)
	gvalid.RegisterRule("did", RuleDidFormat)
}

func RuleETHAddress(ctx context.Context, in gvalid.RuleFuncInput) error {
	eth := in.Value.String()
	if common.IsHexAddress(eth) {
		return nil
	}
	return errors.New(in.Message)
}

func RuleTxHash(ctx context.Context, in gvalid.RuleFuncInput) error {
	tx := in.Value.String()
	if ethutil.IsHexTxHash(tx) {
		return nil
	}
	return errors.New(in.Message)
}

func RuleDidFormat(ctx context.Context, in gvalid.RuleFuncInput) error {
	did := in.Value.String()
	if regutil.IsMetabloxDid(did) {
		return nil
	}
	return errors.New(in.Message)
}
