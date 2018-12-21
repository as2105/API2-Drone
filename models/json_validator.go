package models

import (
	"fmt"

	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
)

type JSONValidationError = gojsonschema.ResultError

type JSONValidator struct {
	schema *gojsonschema.Schema
}

func (v *JSONValidator) Validate(doc []byte) (bool, []JSONValidationError, error) {
	docLoader := gojsonschema.NewBytesLoader(doc)
	result, err := v.schema.Validate(docLoader)
	if err != nil {
		return false, []JSONValidationError{}, errors.Wrap(err, "unable to validate document")
	}
	if result.Valid() {
		return true, []JSONValidationError{}, nil
	}
	return false, result.Errors(), nil
}

func NewJSONValidator(box *packr.Box, resourceType string) (*JSONValidator, error) {
	refName := fmt.Sprintf("file:///fhir.schema.json#/definitions/%s", resourceType)
	sl := gojsonschema.NewSchemaLoader()
	bl := gojsonschema.NewReferenceLoaderFileSystem(refName, box)
	s, err := sl.Compile(bl)
	if err != nil {
		return nil, errors.Wrap(err, "could not compile schema")
	}
	return &JSONValidator{s}, nil
}
