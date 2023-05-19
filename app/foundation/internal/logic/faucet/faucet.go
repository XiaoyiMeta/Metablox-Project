package faucet

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gmlock"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
	"go-metablox/app/foundation/internal/consts"
	"go-metablox/app/foundation/internal/dao"
	"go-metablox/app/foundation/internal/model"
	"go-metablox/app/foundation/internal/model/entity"
	"go-metablox/app/foundation/internal/service"
	"math/big"
	"time"
)

var (
	ErrNotETHAccount            = errors.New("not hex address")
	ErrFaucetIsBusy             = errors.New("faucet is busy now")
	ErrNotRegistered            = errors.New("did not registered")
	ErrBalanceNotEnough         = errors.New("faucet balance not enough")
	ErrFaucetDailyLimitExceeded = errors.New("faucet daily limit exceeded")
	ErrUserAddressLimitExceeded = errors.New("the account has reached the upper limit")
)

type sFaucet struct{}

func init() {
	service.RegisterFaucet(New())
}

func New() *sFaucet {
	return &sFaucet{}
}

func (s *sFaucet) Claims(ctx context.Context, in model.FaucetClaimsInput) (out *model.FaucetClaimsOutput, err error) {

	if !common.IsHexAddress(in.Account) {
		return nil, ErrNotETHAccount
	}
	to := common.HexToAddress(in.Account)
	if err := checkRule(ctx, to); err != nil {
		return nil, err
	}

	if !gmlock.TryLock(consts.LockFaucet) {
		return nil, ErrFaucetIsBusy
	}
	defer gmlock.Unlock(consts.LockFaucet)

	cfg := service.BizCfg().Cfg()

	price, err := GetCacheGasPrice(ctx, service.BizCfg().Client())
	if err != nil {
		return nil, err
	}

	nonce, err := service.BizCfg().Client().PendingNonceAt(ctx, service.BizCfg().FaucetWallet().Address)
	if err != nil {
		return nil, err
	}

	tx := types.NewTx(&types.LegacyTx{
		To:       &to,
		Value:    decimal.NewFromFloat(params.Ether).BigInt(),
		Data:     nil,
		GasPrice: price,
		Gas:      21000,
		Nonce:    nonce,
	})

	rawTx, err := service.BizCfg().WalletStore().SignTxWithPassphrase(service.BizCfg().FaucetWallet(), cfg.FaucetPassphrase, tx, service.BizCfg().ChainID())
	if err != nil {
		return nil, err
	}

	if err := service.BizCfg().Client().SendTransaction(ctx, rawTx); err != nil {
		return nil, err
	}

	dao.DappFaucetRecord.Ctx(ctx).Data(entity.DappFaucetRecord{
		Account: to.Hex(),
		TxHash:  rawTx.Hash().Hex(),
		Amount:  cfg.AmountPerTime,
	}).OmitEmptyData().InsertAndGetId()

	go delayCheckTx(gctx.New(), rawTx.Hash())

	return &model.FaucetClaimsOutput{Hash: rawTx.Hash().Hex()}, err
}

func checkRule(ctx context.Context, to common.Address) error {

	var entity *entity.DappFaucetRecord
	// Is the address has tx pending?
	if err := dao.DappFaucetRecord.Ctx(ctx).
		Where(dao.DappFaucetRecord.Columns().Account, to.Hex()).
		Where(dao.DappFaucetRecord.Columns().TxConfirmed + " is null").
		Limit(1).Scan(&entity); err != nil {
		return err
	}

	if entity != nil {
		if gcron.Search(entity.TxHash) == nil {
			go delayCheckTx(ctx, common.HexToHash(entity.TxHash))
		}
		return gerror.New("please wait last tx committed")
	}

	cfg := service.BizCfg().Cfg()

	// Is the address registered
	registryNonce, err := service.BizCfg().Registry().Nonce(nil, to)
	if err != nil {
		return err
	}
	if registryNonce.Int64() == 0 {
		return ErrNotRegistered
	}
	// Is the balance of treasury sufficient?
	balanceAt, err := service.BizCfg().Client().BalanceAt(ctx, common.HexToAddress(cfg.FaucetAccount), nil)
	if err != nil {
		return err
	}
	if balanceAt.Cmp(decimal.NewFromFloatWithExponent(cfg.AmountPerTime, 18).BigInt()) < 0 {
		return ErrBalanceNotEnough
	}
	// Does the reward exceed the daily limit of treasury?
	dailyLimit, err := dao.DappFaucetRecord.Ctx(ctx).
		Where("date("+dao.DappFaucetRecord.Columns().CreatedAt+") = date(?)", gtime.Now()).
		Where(dao.DappFaucetRecord.Columns().TxConfirmed, true).
		Sum(dao.DappFaucetRecord.Columns().Amount)
	if err != nil {
		return err
	}
	if dailyLimit >= cfg.AmountLimitPerDay {
		return ErrFaucetDailyLimitExceeded
	}
	// Does it exceed the limit of this address?
	addressLimit, err := dao.DappFaucetRecord.Ctx(ctx).
		Where(dao.DappFaucetRecord.Columns().Account, to.Hex()).
		Where(dao.DappFaucetRecord.Columns().TxConfirmed, true).
		Sum(dao.DappFaucetRecord.Columns().Amount)
	if err != nil {
		return err
	}
	if addressLimit >= cfg.AmountLimitPerAddress {
		return ErrUserAddressLimitExceeded
	}
	// Does it meet the minimum claim interval?
	value, err := dao.DappFaucetRecord.Ctx(ctx).
		Where(dao.DappFaucetRecord.Columns().Account, to.Hex()).
		Where(dao.DappFaucetRecord.Columns().TxConfirmed, true).
		Fields(dao.DappFaucetRecord.Columns().CreatedAt).
		OrderDesc(dao.DappFaucetRecord.Columns().CreatedAt).
		Limit(1).Value()
	if err != nil {
		return err
	}
	var latestCreatedAt = value.GTime()
	if latestCreatedAt != nil && gtime.Now().Unix()-latestCreatedAt.Unix() < cfg.ClaimsIntervalHours*3600 {
		return gerror.Newf("claiming too frequently, you can only claim once every %d hours", cfg.ClaimsIntervalHours)
	}

	return nil
}

func GetCacheGasPrice(ctx context.Context, client *ethclient.Client) (*big.Int, error) {
	price, err := gcache.GetOrSetFuncLock(ctx, consts.LockPrice, func(ctx context.Context) (value interface{}, err error) {
		return client.SuggestGasPrice(ctx)
	}, 30*time.Second)
	if err != nil {
		return nil, err
	}
	return big.NewInt(price.Int64()), nil
}

func delayCheckTx(ctx context.Context, hash common.Hash) {
	gcron.DelayAddTimes(ctx, 1*time.Second, "0/3 * * * * *", 5, func(ctx context.Context) {
		data := g.Map{
			dao.DappFaucetRecord.Columns().TxConfirmed: true,
		}
		receipt, err := service.BizCfg().Client().TransactionReceipt(ctx, hash)
		if err != nil && receipt.Status == types.ReceiptStatusSuccessful {
			data[dao.DappFaucetRecord.Columns().Remark] = err.Error()
			return
		}
		data[dao.DappFaucetRecord.Columns().TxConfirmed] = receipt.Status == types.ReceiptStatusSuccessful
		dao.DappFaucetRecord.Ctx(ctx).
			Where(dao.DappFaucetRecord.Columns().TxHash, hash.Hex()).
			Data(data).Update()

		gcron.Remove(hash.Hex())
	}, hash.Hex())
}
