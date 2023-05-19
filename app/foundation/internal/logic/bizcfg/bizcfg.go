package bizcfg

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/MetaBloxIO/did-sdk-go"
	"github.com/MetaBloxIO/did-sdk-go/registry"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"go-metablox/app/foundation/internal/consts"
	"go-metablox/app/foundation/internal/model"
	"go-metablox/app/foundation/internal/service"
	"math/big"
	"os"
	"sync"
)

var (
	ErrWalletAccountsEmpty   = errors.New("the wallet accounts are empty")
	ErrWalletPassphraseEmpty = errors.New("the wallet passphrase is empty")
)
var store = keystore.NewKeyStore("."+keystore.KeyStoreScheme, keystore.StandardScryptN, keystore.StandardScryptP)

type sBizCfg struct {
	sync.Mutex
	model.Config
	client           *ethclient.Client
	chainId          *big.Int
	registryInstance *registry.Registry
}

func init() {
	if len(store.Accounts()) == 0 {
		g.Throw(ErrWalletAccountsEmpty)
	}
	service.RegisterBizCfg(New())
	err := store.Unlock(service.BizCfg().RegistryWallet(), service.BizCfg().Cfg().RegistryPassphrase)
	if err != nil {
		g.Throw(err)
	}
	err = store.Unlock(service.BizCfg().FaucetWallet(), service.BizCfg().Cfg().FaucetPassphrase)
	if err != nil {
		g.Throw(err)
	}
	acc, err := store.Find(service.BizCfg().RegistryWallet())
	if err != nil {
		g.Throw(err)
	}

	err = did.Init(&did.Config{
		Passphrase: service.BizCfg().Cfg().RegistryPassphrase,
		Keystore:   acc.URL.Path,
		ChainId:    service.BizCfg().ChainID(),
	})
	if err != nil {
		g.Throw(err)
	}
}

func keystoreToPrivateKey(privateKeyFile, password string) (*ecdsa.PrivateKey, error) {
	keystoreJSON, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return nil, fmt.Errorf("read keyjson file failed：%s", err.Error())
	}
	key, err := keystore.DecryptKey(keystoreJSON, password)
	if err != nil {
		return nil, err
	}
	//privKey := hex.EncodeToString(unlockedKey.PrivateKey.D.Bytes())
	//addr := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)
	return key.PrivateKey, nil
}

func New() *sBizCfg {
	ctx := gctx.New()
	var cfg sBizCfg
	if g.Cfg().MustGet(ctx, consts.BizCfgKey).Struct(&cfg); &cfg == nil {
		g.Throw(gerror.New("metablox Biz Cfg not found"))
	}
	if err := g.Validator().Data(cfg).Run(ctx); err != nil {
		g.Throw(err)
	}
	if err := cfg.init(ctx); err != nil {
		g.Throw(err)
	}
	return &cfg
}

func (s *sBizCfg) WalletStore() *keystore.KeyStore {
	return store
}

func (s *sBizCfg) init(ctx context.Context) error {
	client, err := ethclient.Dial(s.RpcUrl)
	if err != nil {
		return err
	}
	s.client = client

	chainId, err := client.ChainID(ctx)
	if err != nil {
		return err
	}
	s.chainId = chainId

	registry, err := registry.NewRegistry(common.HexToAddress(s.RegistryAddress), client)
	if err != nil {
		return err
	}
	s.registryInstance = registry

	return nil
}

func (s *sBizCfg) Cfg() model.Config {
	return s.Config
}

func (s *sBizCfg) FaucetWallet() accounts.Account {
	return accounts.Account{Address: common.HexToAddress(s.FaucetAccount)}
}

func (s *sBizCfg) RegistryWallet() accounts.Account {
	return accounts.Account{Address: common.HexToAddress(s.RegistryAccount)}
}

func (s *sBizCfg) Client() *ethclient.Client {
	return s.client
}

func (s *sBizCfg) Registry() *registry.Registry {
	return s.registryInstance
}

func (s *sBizCfg) ChainID() *big.Int {
	return s.chainId
}
