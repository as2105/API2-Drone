package ethereum

import (
	"context"
	"math/big"
	"strings"

	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/config"
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// NewConnection ...
func NewConnection(log *logging.Logger, config *config.Config) (*ethclient.Client, error) {
	return ethclient.Dial(config.RPCURL)
}

// NewTransactOpts ...
func NewTransactOpts(config *config.Config) (*bind.TransactOpts, error) {
	pk := config.PrivateKey
	if pk == "" {
		return nil, errors.New("no private key was provided")
	}
	key, err := crypto.HexToECDSA(strings.TrimPrefix(pk, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse private key")
	}
	t := bind.NewKeyedTransactor(key)
	t.Context = context.Background()
	t.GasPrice = big.NewInt(config.GasPrice)
	t.GasLimit = config.GasLimit
	return t, nil
}
