package ethereum

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/storage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/vincent-petithory/dataurl"
)

const FHIRJSONMediaType = "application/fhir+json"

type ObjectCollectionElementBytesData struct {
	data      []byte
	mediaType string
}

func (o *ObjectCollectionElementBytesData) URI() string {
	uri := dataurl.New(o.data, o.mediaType)
	return uri.String()
}

func NewObjectCollectionElementFHIRJSONData(data []byte) *ObjectCollectionElementBytesData {
	return &ObjectCollectionElementBytesData{
		data:      data,
		mediaType: FHIRJSONMediaType,
	}
}

type ObjectCollectionElementIPFSData struct {
	address string
	path    string
}

func (o *ObjectCollectionElementIPFSData) URI() string {
	s := fmt.Sprintf("ipfs:%s", o.address)
	if o.path != "" {
		s = fmt.Sprintf("%s/%s", s, strings.TrimPrefix(o.path, "/"))
	}
	return s
}

type ObjectCollectionElementURIData struct {
	uri string
}

func (o *ObjectCollectionElementURIData) URI() string {
	return o.uri
}

type ObjectCollectionElement struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      storage.ObjectCollectionElementData
}

func NewObjectCollectionElement(uri string, createdAt, updatedAt *big.Int) (*ObjectCollectionElement, error) {
	newObj := &ObjectCollectionElement{
		CreatedAt: bigintToTime(createdAt),
		UpdatedAt: bigintToTime(updatedAt),
	}

	dataElems := strings.Split(uri, ":")
	switch dataElems[0] {
	case "ipfs":
		ipfsData := strings.Split(dataElems[1], "/")
		newObj.Data = &ObjectCollectionElementIPFSData{ipfsData[0], ipfsData[1]}
	case "data":
		u, err := dataurl.DecodeString(uri)
		if err != nil {
			return nil, errors.Wrap(err, "unable to decode data uri")
		}
		switch u.MediaType.ContentType() {
		case FHIRJSONMediaType:
			newObj.Data = NewObjectCollectionElementFHIRJSONData(u.Data)
		default:
			newObj.Data = &ObjectCollectionElementBytesData{u.Data, u.MediaType.ContentType()}
		}
	default:
		newObj.Data = &ObjectCollectionElementURIData{uri}
	}

	return newObj, nil
}

type ObjectIndexKey = [32]byte

type ObjectIndexPopulationFunc func(storage.ObjectCollectionElementData) (ObjectIndexKey, error)

type ObjectIndexCollection map[common.Address]ObjectIndexPopulationFunc
