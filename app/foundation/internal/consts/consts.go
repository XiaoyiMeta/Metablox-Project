package consts

import (
	"github.com/ethereum/go-ethereum/common"
	"go-metablox/pkg/biz"
)

const AppName = "foundation"
const AppUsage = "foundation service"

const (
	BizCtxKey     = "BizCtx"
	BizCtxUserKey = "BizCtxUser"
)

var ZeroAddress = common.Address{}

const LockFaucet = "lock:faucet"
const LockPrice = "lock:price"

const BizCfgKey = biz.ProjectName + "." + AppName

type VCType string

const (
	Validator VCType = "Validator"
)

var vcTypeMap = map[VCType]string{
	Validator: "Wifi Access Credential",
}

func (c VCType) Description() string {
	return vcTypeMap[c]
}

func (c VCType) Type() string {
	return string(c)
}
