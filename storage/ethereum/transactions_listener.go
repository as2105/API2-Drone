package ethereum

import (
	"context"
	"github.com/SynapticHealthAlliance/fhir-api/config"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/fx"
	"time"
)

// TransactionsChannel ...
type TransactionsChannel chan *types.Transaction

// NewTransactionsChannel ...
func NewTransactionsChannel(c *config.Config) TransactionsChannel {
	return make(TransactionsChannel, c.TransactionsChannelBuffer)
}

// TransactionsListener ...
type TransactionsListener struct {
	log           logging.FieldLogger
	txnsChan      TransactionsChannel
	context       context.Context
	contextCancel context.CancelFunc
	connection    *ethclient.Client
}

// Start ...
func (l *TransactionsListener) Start() {
	l.log.Info("started")
	go func() {
		for {
			select {
			case txn := <-l.txnsChan:
				l.processTxn(txn)
			case <-l.context.Done():
				l.log.Info("stopped")
				return
			}
		}
	}()
}

func (l *TransactionsListener) processTxn(txn *types.Transaction) {
	log := l.log.WithField("txn", txn.Hash().String())
	log.Debug("listener received transaction")
	ctx, cancel := context.WithTimeout(l.context, (30 * time.Second))
	defer cancel()
	receipt, err := bind.WaitMined(ctx, l.connection, txn)
	if err != nil {
		log.WithError(err).Error("error reported while waiting for transaction to be mined")
	}
	if receipt != nil {
		msg := "receipt received"
		entry := log.WithFields(logging.Fields{
			"gas_used":       receipt.GasUsed,
			"total_gas_used": receipt.CumulativeGasUsed,
		})
		if receipt.Status != types.ReceiptStatusSuccessful {
			entry.WithField("status", "failed").Error(msg)
			return
		}
		entry.WithField("status", "success").Info(msg)
	}
}

// NewTransactionsListener ...
func NewTransactionsListener(lc fx.Lifecycle, log *logging.Logger, txnsChan TransactionsChannel, conn *ethclient.Client) *TransactionsListener {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	txnL := &TransactionsListener{
		log:           log.WithField("component", "listener"),
		txnsChan:      txnsChan,
		context:       ctx,
		contextCancel: cancel,
		connection:    conn,
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			cancel()
			return nil
		},
	})

	return txnL
}

// StartTransactionsListener ...
func StartTransactionsListener(listener *TransactionsListener) {
	listener.Start()
}
