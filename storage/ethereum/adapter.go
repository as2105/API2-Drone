package ethereum

import (
	"fmt"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/storage"
	"github.com/SynapticHealthAlliance/pdx-contracts/go/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

type EthereumAdapter struct {
	connection              bind.ContractBackend
	transactor              bind.ContractTransactor
	objectCollectionAddress common.Address
	organizationAddress     common.Address
	objectIndexCollection   ObjectIndexCollection

	objectCollectionCaller     contracts.ObjectCollectionCaller
	objectCollectionTransactor contracts.ObjectCollectionTransactor

	submittedTransactions chan<- *types.Transaction
}

func (a *EthereumAdapter) Create(data storage.ObjectCollectionElementData) (uuid.UUID, error) {
	newUUID := uuid.NewUUID()
	oidxAddrs := make([]common.Address, len(a.objectIndexCollection))
	oidxKeys := make([][32]byte, len(a.objectIndexCollection))
	for oidxAddr, oidxFunc := range a.objectIndexCollection {
		oidxAddrs = append(oidxAddrs, oidxAddr)
		newKey, err := oidxFunc(data)
		if err != nil {
			return newUUID, errors.Wrap(err, fmt.Sprintf("failed to create key for index at address %q", oidxAddr.String()))
		}
		oidxKeys = append(oidxKeys, newKey)
	}
	txn, err := a.objectCollectionTransactor.AddObject(nil, newUUID.Array(), a.organizationAddress, data.URI(), oidxAddrs, oidxKeys)
	return newUUID, a.handleTransaction(txn, err)
}

func (a *EthereumAdapter) handleTransaction(txn *types.Transaction, err error) error {
	if err == nil {
		a.submittedTransactions <- txn
	}
	return err
}

func (a *EthereumAdapter) Update(id uuid.UUID, lastUpdatedAt time.Time, data storage.ObjectCollectionElementData, changeScore uint8) error {
	return a.handleTransaction(
		a.objectCollectionTransactor.UpdateObject(nil, id.Array(), timeToBigint(lastUpdatedAt), data.URI(), changeScore),
	)
}

func (a *EthereumAdapter) Destroy(id uuid.UUID) error {
	return a.handleTransaction(
		a.objectCollectionTransactor.RemoveObject(nil, id.Array()),
	)
}

func (a *EthereumAdapter) Read(id uuid.UUID) (*ObjectCollectionElement, error) {
	r, err := a.objectCollectionCaller.GetObject(nil, id.Array())
	if err != nil {
		return nil, errors.Wrap(err, "unable to read from contract")
	}
	return NewObjectCollectionElement(r.Uri, r.CreatedAt, r.UpdatedAt)
}

func NewEthereumAdapter(
	connection bind.ContractBackend,
	transactor bind.ContractTransactor,
	objectCollectionAddress common.Address,
	organizationAddress common.Address,
	objectIndexCollection ObjectIndexCollection,
	transactionBuffer uint,
) (*EthereumAdapter, error) {
	coll, err := contracts.NewObjectCollection(objectCollectionAddress, connection)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get collection contract")
	}
	return &EthereumAdapter{
		connection:                 connection,
		transactor:                 transactor,
		objectCollectionAddress:    objectCollectionAddress,
		organizationAddress:        organizationAddress,
		objectIndexCollection:      objectIndexCollection,
		objectCollectionCaller:     coll.ObjectCollectionCaller,
		objectCollectionTransactor: coll.ObjectCollectionTransactor,
		submittedTransactions:      make(chan *types.Transaction, transactionBuffer),
	}, nil
}
