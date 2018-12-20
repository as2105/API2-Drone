// Copyright © 2018 Optum

package storage

import (
	"time"

	"github.com/pborman/uuid"
)

type ObjectCollectionElementData interface {
	URI() string
}

type StorageAdapter interface {
	Create(data ObjectCollectionElementData) (uuid.UUID, error)
	Update(id uuid.UUID, lastUpdatedAt time.Time, data ObjectCollectionElementData, changeScore uint8) error
	Destroy(id uuid.UUID) error
	Read(id uuid.UUID) (ObjectCollectionElementData, error)
}
