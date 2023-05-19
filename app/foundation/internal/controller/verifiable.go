package controller

import (
	"context"
	"github.com/MetaBloxIO/did-sdk-go"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gmlock"
	v1 "go-metablox/app/foundation/api/v1"
	"go-metablox/app/foundation/internal/consts"
	"go-metablox/app/foundation/internal/model"
	"go-metablox/app/foundation/internal/service"
	"go-metablox/pkg/biz"
	"go-metablox/pkg/gconsts"
	"go-metablox/pkg/util/ethutil"
)

var Verifiable = &cVerifiable{}

type cVerifiable struct {
}

func (c *cVerifiable) Nonce(ctx context.Context, req *v1.VCNonceReq) (*v1.VCNonceRes, error) {
	out, err := service.Verifiable().Nonce(ctx, model.VCNonceInput{
		Did: req.Did,
	})
	return &v1.VCNonceRes{Nonce: out.Nonce, ExpiresAt: out.ExpiresAt}, err
}

func (c *cVerifiable) VerifyVP(ctx context.Context, req *v1.VerifyVPReq) (*v1.VerifyVPRes, error) {
	service.Verifiable().VerifyVP(ctx, req.VerifiablePresentation)
	return &v1.VerifyVPRes{}, nil
}

func (c *cVerifiable) IssueWifiVC(ctx context.Context, req *v1.IssueWifiVCReq) (*v1.IssueWifiVCRes, error) {
	if err := checkNonce(ctx, req.Nonce, req.Id); err != nil {
		return nil, err
	}

	prepareDID, _ := did.PrepareDID(req.Id)
	address, err := service.BizCfg().Registry().Dids(nil, prepareDID[2])
	if err != nil {
		return nil, err
	}
	if address == gconsts.ZeroAddress {
		return nil, biz.NewError("did not registered")
	}

	recAddress, err := ethutil.EcRecover([]byte(req.Nonce), hexutil.MustDecode(req.Signature))
	if recAddress != address {
		return nil, gerror.NewCode(biz.CodeError, "signature verification failed")
	}
	if !gmlock.TryLock(consts.LockPrefixVC + req.Id) {
		return nil, biz.NewError("operation too frequent, try again later")
	}
	defer gmlock.Unlock(consts.LockPrefixVC + req.Id)

	vc, err := service.Verifiable().IssueVC(ctx, model.IssueVCInput{
		Id:   req.Id,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &v1.IssueWifiVCRes{
		VerifiableCredential: vc.VerifiableCredential,
	}, nil
}

func (c *cVerifiable) RenewWifiVC(ctx context.Context, req *v1.RenewWifiVCReq) (*v1.RenewWifiVCRes, error) {
	if err := checkNonce(ctx, req.Proof.Nonce, req.Holder); err != nil {
		return nil, err
	}

	vc, err := service.Verifiable().RenewVC(ctx, model.VCInput{
		VerifiablePresentation: req.VerifiablePresentation,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RenewWifiVCRes{
		VerifiableCredential: vc.VerifiableCredential,
	}, nil
}

func (c *cVerifiable) RevokeWifiVC(ctx context.Context, req *v1.RevokeWifiVCReq) (*v1.RevokeWifiVCRes, error) {
	if err := checkNonce(ctx, req.Proof.Nonce, req.Holder); err != nil {
		return nil, err
	}

	return &v1.RevokeWifiVCRes{}, nil
}

func (c *cVerifiable) IssueMiningVC(ctx context.Context, req *v1.IssueMiningVCReq) (*v1.IssueMiningVCRes, error) {
	if err := checkNonce(ctx, req.Nonce, req.Id); err != nil {
		return nil, err
	}

	prepareDID, _ := did.PrepareDID(req.Id)
	address, err := service.BizCfg().Registry().Dids(nil, prepareDID[2])
	if err != nil {
		return nil, err
	}
	if address == gconsts.ZeroAddress {
		return nil, biz.NewError("did not registered")
	}

	recAddress, err := ethutil.EcRecover([]byte(req.Nonce), hexutil.MustDecode(req.Signature))
	if recAddress != address {
		return nil, gerror.NewCode(biz.CodeError, "signature verification failed")
	}
	if !gmlock.TryLock(consts.LockPrefixVC + req.Id) {
		return nil, biz.NewError("operation too frequent, try again later")
	}
	defer gmlock.Unlock(consts.LockPrefixVC + req.Id)

	return &v1.IssueMiningVCRes{}, nil
}

func (c *cVerifiable) RenewMiningVC(ctx context.Context, req *v1.RenewMiningVCReq) (*v1.RenewMiningVCRes, error) {
	if err := checkNonce(ctx, req.Proof.Nonce, req.Holder); err != nil {
		return nil, err
	}
	return &v1.RenewMiningVCRes{}, nil
}

func (c *cVerifiable) RevokeMiningVC(ctx context.Context, req *v1.RevokeMiningVCReq) (*v1.RevokeMiningVCRes, error) {
	if err := checkNonce(ctx, req.Proof.Nonce, req.Holder); err != nil {
		return nil, err
	}
	return &v1.RevokeMiningVCRes{}, nil
}

func checkNonce(ctx context.Context, nonce string, did string) error {
	value, _ := gcache.Get(ctx, consts.CachePrefixChallenge+nonce)
	if value.IsEmpty() {
		return biz.NewError("challenge timeout or not exists")
	}
	if value.String() != did {
		return biz.NewError("wrong did")
	}
	defer gcache.Remove(ctx, consts.CachePrefixChallenge+nonce)
	return nil
}
