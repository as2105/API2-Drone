// Copyright © 2018 Optum

package handlers

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/handlers/resources"
	"github.com/SynapticHealthAlliance/fhir-api/handlers/responders"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/utils"
	"github.com/gorilla/mux"
)

// RegisterFHIRCapabilityStatementRoutes mounts our HTTP handler on the mux.
func RegisterFHIRCapabilityStatementRoutes(r *mux.Router, log *logging.Logger, cConfig *resources.CapabilityConfig) {
	log.Debug("executing RegisterFHIRCapabilityStatementRoutes")
	h := responders.JSON(cConfig, http.StatusOK)
	r.Handle("/", h).Methods("GET")
	r.Handle("/metadata", h).Methods("GET")
}

// RegisterAllFHIRResourceRoutes ...
func RegisterAllFHIRResourceRoutes(r *mux.Router, log *logging.Logger) {
	log.Debug("executing RegisterAllFHIRResourceRoutes")
	for _, i := range resources.GetRegisteredResources() {
		registerFHIRResourceRoutes(r, log, i)
	}
}

func registerFHIRResourceRoutes(r *mux.Router, log *logging.Logger, i interface{}) {
	dLog := log.WithField("function", "registerFHIRResourceRoutes")
	dLog.Debug("executing")

	seg := utils.GetBaseTypeName(i)
	dLog = dLog.WithField("resource", seg)

	typePrefix := "/" + seg
	instancePrefix := typePrefix + "/{resourceID}"

	if t, ok := i.(resources.CreateableResource); ok {
		dLog.Debug("registering create method")
		r.Handle(typePrefix, t.Create(log)).Methods("POST")
	}

	if t, ok := i.(resources.SearchableResource); ok {
		dLog.Debug("registering search method")
		r.Handle(typePrefix, t.Search(log)).Methods("GET")
	}

	if t, ok := i.(resources.UpdateableResource); ok {
		dLog.Debug("registering update method")
		r.Handle(instancePrefix, t.Update(log)).Methods("PUT")
	}

	if t, ok := i.(resources.DeleteableResource); ok {
		dLog.Debug("registering delete method")
		r.Handle(instancePrefix, t.Delete(log)).Methods("DELETE")
	}

	if t, ok := i.(resources.ReadableResource); ok {
		dLog.Debug("registering read method")
		r.Handle(instancePrefix, t.Read(log)).Methods("GET")
	}

	if t, ok := i.(resources.PatchableResource); ok {
		dLog.Debug("registering patch method")
		r.Handle(instancePrefix, t.Patch(log)).Methods("PATCH")
	}

	if t, ok := i.(resources.VersionReadableResource); ok {
		dLog.Debug("registering version read method")
		r.Handle(instancePrefix+"/_history/{versionID}", t.VersionRead(log)).Methods("GET")
	}

	if t, ok := i.(resources.ValidateableResource); ok {
		dLog.Debug("registering validate method")
		r.Handle(typePrefix+"/$validate", t.Validate(log)).Methods("POST")
	}
}
