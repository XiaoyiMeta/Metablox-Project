package openroaming

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"go-metablox/app/foundation/internal/dao"
	"go-metablox/app/foundation/internal/model"
	"go-metablox/app/foundation/internal/model/do"
	"go-metablox/app/foundation/internal/model/entity"
	"go-metablox/app/foundation/internal/service"
)

const Radius = "radius"
const DomainSuffix = "@metablox.io"

type sOpenroaming struct {
}

func init() {
	service.RegisterOpenroaming(New())
}

func New() *sOpenroaming {
	return &sOpenroaming{}
}

func (o *sOpenroaming) GetUserInfo(ctx context.Context, username string) (*model.WifiUserInfoOutput, error) {
	var one *entity.Radcheck

	if err := g.DB(Radius).Model(do.Radcheck{}).Ctx(ctx).Where(dao.Radcheck.Columns().Username, username).
		Scan(&one); err != nil {
		return nil, err
	}

	if one != nil {
		return &model.WifiUserInfoOutput{
			Username: one.Username + DomainSuffix,
			Password: one.Value,
		}, nil
	}
	// Automatically create a new user if not exists
	password := guid.S()
	if _, err := o.CreateUser(ctx, model.CreateWifiUserInfoInput{
		Username: username,
		Password: password,
	}); err != nil {
		return nil, err
	}

	return &model.WifiUserInfoOutput{
		Username: username + DomainSuffix,
		Password: password,
	}, nil
}

func (o *sOpenroaming) CreateUser(ctx context.Context, in model.CreateWifiUserInfoInput) (int64, error) {

	return g.DB(Radius).Ctx(ctx).Model(do.Radcheck{}).Data(do.Radcheck{
		Username:  in.Username,
		Attribute: "Cleartext-Password",
		Op:        ":=",
		Value:     in.Password,
	}).InsertAndGetId()

}
