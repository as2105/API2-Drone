package ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SynapticHealthAlliance/fhir-api/config"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/pdx-contracts/go/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"time"
)

// Adapter ...
type Adapter struct {
	connection                 *ethclient.Client
	transactOpts               *bind.TransactOpts
	organizationAddress        common.Address
	objectCollectionContract   *config.ObjectCollectionContract
	objectCollectionCaller     contracts.ObjectCollectionCaller
	objectCollectionTransactor contracts.ObjectCollectionTransactor
	submittedTransactions      chan<- *types.Transaction
	log                        logging.FieldLogger
}

// Create ...
func (a *Adapter) Create(ctx context.Context, uuid uuid.UUID, data ObjectCollectionElementData) error {
	log := a.log.WithField("uuid", uuid.String())
	addrs := []common.Address{}
	keys := []objectIndexKey{}
	log.WithField("uri", data.URI()).Debug("storing data")
	bytes, err := data.Bytes()
	if err != nil {
		return errors.Wrap(err, "failed to get bytes from data")
	}
	var jsonData interface{}
	log.Debugf("unmarshalling JSON data: %v", string(bytes))
	err = json.Unmarshal(bytes, &jsonData)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal json data")
	}
	log.Debug("generating index keys")
	for _, idx := range a.objectCollectionContract.Indexes {
		// it is possible that a json path could return multiple results, so for each result we will add the key and the address of the index
		log.Debugf("generating index key for address %v using %v", idx.Address.String(), idx.JSONPath.String())
		newKeys, err := a.generateObjectIndexKeys(idx, jsonData)
		if err != nil {
			return errors.Wrapf(err, "unable to generate index key from data using provided json path %q", idx.JSONPath)
		}
		keys = append(keys, newKeys...)
		for i := 0; i < len(newKeys); i++ {
			addrs = append(addrs, idx.Address)
		}
	}
	log.Debugf("index keys (%d): %v", len(keys), keys)
	log.Debugf("index addresses (%d): %v", len(addrs), addrs)
	txn, err := a.objectCollectionTransactor.AddObject(a.transactOpts, uuid.Array(), a.organizationAddress, data.URI(), addrs, keys)
	return a.handleTransaction(txn, err)
}

func (a *Adapter) handleTransaction(txn *types.Transaction, err error) error {
	if txn != nil {
		a.log.Debugf("transaction received: %v", txn.Hash().String())
		if err == nil {
			a.submittedTransactions <- txn
		}
	}
	return err
}

// Update ...
func (a *Adapter) Update(ctx context.Context, id uuid.UUID, lastUpdatedAt time.Time, data ObjectCollectionElementData, changeScore uint8) error {
	return a.handleTransaction(
		a.objectCollectionTransactor.UpdateObject(a.transactOpts, id.Array(), timeToBigint(lastUpdatedAt), data.URI(), changeScore),
	)
}

// Destroy ...
func (a *Adapter) Destroy(ctx context.Context, id uuid.UUID) error {
	return a.handleTransaction(
		a.objectCollectionTransactor.RemoveObject(a.transactOpts, id.Array()),
	)
}

// Read ...
func (a *Adapter) Read(ctx context.Context, id uuid.UUID) (*ObjectCollectionElement, error) {
	r, err := a.objectCollectionCaller.GetObject(nil, id.Array())
	if err != nil {
		return nil, errors.Wrap(err, "unable to read from contract")
	}
	return NewObjectCollectionElement(r.Uri, r.CreatedAt, r.UpdatedAt)
}

// ReadJSONResource ...
func (a *Adapter) ReadJSONResource(ctx context.Context, id uuid.UUID, resource interface{}) error {
	data, err := a.Read(ctx, id)
	if err != nil {
		return errors.Wrap(err, "failed to find record")
	}
	jsonBytes, err := data.Data.Bytes()
	if err != nil {
		return errors.Wrap(err, "failed to bytes from object data")
	}
	if err := json.Unmarshal(jsonBytes, resource); err != nil {
		return errors.Wrap(err, "failed to unmarshall JSON bytes")
	}
	return nil
}

// key creates a 32-byte index key from a JSON document using a JSON path
func (a *Adapter) generateObjectIndexKeys(idx *config.ObjectIndex, jsonData interface{}) ([]objectIndexKey, error) {
	keys := []objectIndexKey{}
	result, err := idx.JSONPath.Lookup(jsonData)
	if err != nil {
		return keys, errors.Wrap(err, "failed to create key from provided JSON path")
	}
	a.log.Debugf("lookup result: %v (%v)", result, jsonData)
	var rawKeys []interface{}
	switch t := result.(type) {
	case []interface{}:
		rawKeys = t
	default:
		rawKeys = []interface{}{t}
	}
	for _, k := range rawKeys {
		keyStr := fmt.Sprintf("%v", k)
		newKey, err := a.stringToIndexKey(keyStr)
		if err != nil {
			return keys, errors.Errorf("failed to generate key from string %q", keyStr)
		}
		keys = append(keys, newKey)
	}
	return keys, nil
}

func (a *Adapter) stringToIndexKey(str string) (objectIndexKey, error) {
	bytes := []byte(str)
	var key objectIndexKey
	if len(bytes) > 32 {
		return key, errors.Errorf("string provided is too long; must be less than 32 characters: %q", str)
	}
	copy(key[:], bytes)
	return key, nil
}

// NewAdapter ...
func NewAdapter(
	connection *ethclient.Client,
	transactOpts *bind.TransactOpts,
	organizationAddress common.Address,
	objectCollectionContract *config.ObjectCollectionContract,
	submittedTransactions chan<- *types.Transaction,
	log *logging.Logger,
) (*Adapter, error) {
	coll, err := contracts.NewObjectCollection(objectCollectionContract.Address, connection)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get collection contract")
	}
	return &Adapter{
		connection:                 connection,
		transactOpts:               transactOpts,
		organizationAddress:        organizationAddress,
		objectCollectionContract:   objectCollectionContract,
		objectCollectionCaller:     coll.ObjectCollectionCaller,
		objectCollectionTransactor: coll.ObjectCollectionTransactor,
		submittedTransactions:      submittedTransactions,
		log:                        log.WithField("component", "ethereum"),
	}, nil
}
