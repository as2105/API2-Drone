// Copyright © 2018 Optum
package handlers

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/internal/persisteth"
)

type Logger interface {
	Fatal(error)
}

type Validator interface {
	Validate(interface{}) error
}

func NewCreateHandler(store persisteth.ResourceStore, logger Logger, validator Validator) http.HandlerFunc {

	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		var obj interface{}

		validator.Validate(obj)

		uuid, etag, err := store.Create(obj)
		if err != nil {
			logger.Fatal(err)
		}

		_ = uuid
		_ = etag
	})

}
