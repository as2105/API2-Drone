package ethereum

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/vincent-petithory/dataurl"
)

const fhirJSONMediaType = "application/fhir+json"

// ObjectCollectionElementData ...
type ObjectCollectionElementData interface {
	URI() string
	Bytes() ([]byte, error)
}

// ObjectCollectionElementBytesData ...
type ObjectCollectionElementBytesData struct {
	data      []byte
	mediaType string
}

// URI ...
func (o *ObjectCollectionElementBytesData) URI() string {
	uri := dataurl.New(o.data, o.mediaType)
	return uri.String()
}

// Bytes ...
func (o *ObjectCollectionElementBytesData) Bytes() ([]byte, error) {
	return o.data, nil
}

// NewObjectCollectionElementFHIRJSONData ...
func NewObjectCollectionElementFHIRJSONData(data []byte) *ObjectCollectionElementBytesData {
	return &ObjectCollectionElementBytesData{
		data:      data,
		mediaType: fhirJSONMediaType,
	}
}

// ObjectCollectionElementIPFSData ...
type ObjectCollectionElementIPFSData struct {
	address string
	path    string
}

// URI ...
func (o *ObjectCollectionElementIPFSData) URI() string {
	s := fmt.Sprintf("ipfs:%s", o.address)
	if o.path != "" {
		s = fmt.Sprintf("%s/%s", s, strings.TrimPrefix(o.path, "/"))
	}
	return s
}

// Bytes ...
func (o *ObjectCollectionElementIPFSData) Bytes() ([]byte, error) {
	// TODO: fetch from IPFS
	return []byte{}, nil
}

// ObjectCollectionElementURIData ...
type ObjectCollectionElementURIData struct {
	uri string
}

// URI ...
func (o *ObjectCollectionElementURIData) URI() string {
	return o.uri
}

// Bytes ...
func (o *ObjectCollectionElementURIData) Bytes() ([]byte, error) {
	// TODO: fetch?
	return []byte{}, nil
}

// ObjectCollectionElement ...
type ObjectCollectionElement struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      ObjectCollectionElementData
}

// NewObjectCollectionElement ...
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
		case fhirJSONMediaType:
			newObj.Data = NewObjectCollectionElementFHIRJSONData(u.Data)
		default:
			newObj.Data = &ObjectCollectionElementBytesData{u.Data, u.MediaType.ContentType()}
		}
	default:
		newObj.Data = &ObjectCollectionElementURIData{uri}
	}

	return newObj, nil
}

type objectIndexKey = [32]byte
