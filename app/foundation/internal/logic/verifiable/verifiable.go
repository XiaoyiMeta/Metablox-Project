package verifiable

import (
	"context"
	"crypto/sha256"
	"github.com/MetaBloxIO/did-sdk-go"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"go-metablox/app/foundation/internal/consts"
	"go-metablox/app/foundation/internal/dao"
	"go-metablox/app/foundation/internal/model"
	"go-metablox/app/foundation/internal/model/do"
	"go-metablox/app/foundation/internal/model/entity"
	"go-metablox/app/foundation/internal/service"
	"go-metablox/pkg/biz"
	"strconv"
	"time"
)

type sVerifiable struct {
}

func init() {
	service.RegisterVerifiable(New())
}

func New() *sVerifiable {
	return &sVerifiable{}
}

func (s *sVerifiable) Nonce(ctx context.Context, in model.VCNonceInput) (*model.VCChallengeOutput, error) {

	nonce := guid.S()
	if err := gcache.Set(ctx, consts.CachePrefixChallenge+nonce, in.Did, 5*time.Minute); err != nil {
		return nil, err
	}
	g.Log().Infof(ctx, "did: %s, nonce: %s,nonce generated!", in.Did, nonce)
	return &model.VCChallengeOutput{
		ExpiresAt: time.Now().Add(5 * time.Minute).Format(time.RFC3339),
		Nonce:     nonce,
	}, nil
}

func (s *sVerifiable) VerifyVP(ctx context.Context, in did.VerifiablePresentation) error {
	vp, err := did.VerifyVP(&in, service.BizCfg().Registry())
	if err != nil {
		return err
	}
	if !vp {
		return biz.NewError("vp verify failed")
	}
	return nil
}

func (s *sVerifiable) VerifyVC(ctx context.Context, in did.VerifiableCredential) error {

	vp, err := did.VerifyVC(&in, service.BizCfg().Registry())
	if err != nil {
		return err
	}
	if !vp {
		return biz.NewError("vc verify failed")
	}
	return nil
}

func (s *sVerifiable) IssueVC(ctx context.Context, in model.IssueVCInput) (*model.IssueVCOutput, error) {

	issuerDocument, err := getDidDocument(did.GetIssuerDid())
	if err != nil {
		g.Log().Warningf(ctx, "get vc issuer document failed: %v", err)
		return nil, biz.NewError("get vc issuer document failed")
	}

	var vc *entity.DidVerifiableCredential
	err = dao.DidVerifiableCredential.Ctx(ctx).
		Where(dao.DidVerifiableCredential.Columns().Holder, in.Id).
		Where(dao.DidVerifiableCredential.Columns().Type, in.Type).
		OrderDesc(dao.DidVerifiableCredential.Columns().Id).
		Limit(1).
		Scan(&vc)
	if err != nil {
		return nil, err
	}

	var newVC *did.VerifiableCredential
	if vc == nil {
		newVC, err = createWifiVC(ctx, issuerDocument, in.Type, in.Id, g.Map{
			"id":   in.Id,
			"type": in.Type,
		})
		if err != nil {
			return nil, err
		}

		return &model.IssueVCOutput{VerifiableCredential: newVC}, nil
	}

	if !vc.RevokedAt.IsZero() {
		return nil, biz.NewError("vc has been revoked")
	}

	if vc.ExpirationDate.Before(time.Now()) {
		if err = gconv.Scan(vc, &newVC); err != nil {
			return nil, err
		}
		return &model.IssueVCOutput{VerifiableCredential: newVC}, nil
	}

	if err = gconv.Scan(vc.RawText, &newVC); err != nil {
		return nil, err
	}

	return &model.IssueVCOutput{VerifiableCredential: newVC}, nil
}

func (s *sVerifiable) RenewVC(ctx context.Context, in model.VCInput) (*model.VCOutput, error) {
	err := s.VerifyVP(ctx, in.VerifiablePresentation)
	if err != nil {
		return nil, err
	}

	for _, vc := range in.VerifiableCredential {
		err = s.VerifyVC(ctx, vc)
		if err != nil {
			return nil, err
		}
	}

	return &model.VCOutput{}, nil
}

func (s *sVerifiable) RevokeVC(ctx context.Context, in model.VCInput) error {

	return nil
}

func getDidDocument(id string) (*did.DIDDocument, error) {
	opts := &did.ResolutionOptions{}
	resolutionMeta, issuerDocument, _ := did.Resolve(id, opts, service.BizCfg().Registry())
	if resolutionMeta.Error != "" {
		return nil, biz.NewError(resolutionMeta.Error)
	}

	return issuerDocument, nil

}

func createWifiVC(ctx context.Context, issuer *did.DIDDocument, vcType string, holderDid string, vcData interface{}) (*did.VerifiableCredential, error) {

	vc, err := did.CreateVC(issuer)
	if err != nil {
		return nil, err
	}

	vc.Type = append(vc.Type, did.TypeWifi)
	vc.Description = "Wifi Access Credential"
	vc.CredentialSubject = vcData

	//Upload VC to DB and generate ID. Has to be done before creating signature, as changing the ID will change the signature
	err = did.ConvertTimesToDBFormat(vc)
	if err != nil {
		return nil, err
	}

	err = did.ConvertTimesFromDBFormat(vc)
	if err != nil {
		return nil, err
	}

	vcStore := &do.DidVerifiableCredential{
		ChainId:           service.BizCfg().ChainID().Int64(),
		Type:              vcType,
		Context:           gconv.String(vc.Context),
		Issuer:            did.GetIssuerDid(),
		Holder:            holderDid,
		CredentialSubject: gconv.String(vc.CredentialSubject),
		Description:       vc.Description,
		Proof:             gconv.String(vc.Proof),
		IssuanceDate:      time.Now(),
		ExpirationDate:    time.Now().AddDate(1, 0, 0),
	}

	id, err := dao.DidVerifiableCredential.Ctx(ctx).InsertAndGetId(vcStore)
	if err != nil {
		return nil, err
	}
	vc.ID = did.BaseIDString + strconv.Itoa(int(id))
	_, err = dao.DidVerifiableCredential.Ctx(ctx).WherePri(id).
		Data(dao.DidVerifiableCredential.Columns().RawText, gconv.String(vc)).
		Update()
	if err != nil {
		return nil, err
	}

	//Create the proof's signature using a stringified version of the VC and the issuer's private key.
	//This way, the signature can be verified by re-stringifying the VC and looking up the public key in the issuer's DID document.
	//Verification will only succeed if the VC was unchanged since the signature and if the issuer
	//public key matches the private key used to make the signature
	hashedVC := sha256.Sum256(did.ConvertVCToBytes(*vc))

	signatureData, err := did.CreateJWSSignature(did.GetIssuerPrivateKey(), hashedVC[:])
	if err != nil {
		return nil, err
	}
	vc.Proof.JWSSignature = signatureData

	return vc, nil
}
