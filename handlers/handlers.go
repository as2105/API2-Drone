// Copyright © 2018 Optum
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// RegisterFHIRCapabilityStatement mounts our HTTP handler on the mux.
func RegisterFHIRCapabilityStatement(r *mux.Router, logger *logrus.Logger) {
	r.Handle("/", newCapabilityStatementHandler(logger)).Methods("GET")
	r.Handle("/metadata", newCapabilityStatementHandler(logger)).Methods("GET")
}

func newCapabilityStatementHandler(logger *logrus.Logger) http.Handler {
	logger.Print("Executing newCapabilityStatementHandler.")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		v := map[string]string{
			"message": "hello world",
		}
		json.NewEncoder(rw).Encode(v)
	})
}
