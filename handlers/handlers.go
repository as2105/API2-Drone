// Copyright © 2018 Optum

package handlers

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/handlers/resources"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/utils"
	"github.com/gorilla/mux"
	"github.com/heptiolabs/healthcheck"
	"github.com/unrolled/render"
)

// RegisterFHIRCapabilityStatementRoutes mounts our HTTP handler on the mux.
func RegisterFHIRCapabilityStatementRoutes(r *mux.Router, log *logging.Logger, cConfig *resources.CapabilityConfig, rndr *render.Render) {
	log.Debug("executing RegisterFHIRCapabilityStatementRoutes")
	h := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Last-Modified", cConfig.Date)
		rndr.JSON(w, http.StatusOK, cConfig)
	})
	r.Handle("/", h).Methods("GET")
	r.Handle("/metadata", h).Methods("GET")
}

// RegisterAllFHIRResourceRoutes ...
func RegisterAllFHIRResourceRoutes(r *mux.Router, log *logging.Logger, rndr *render.Render, registry resources.ResourceRegistry) {
	log.Debug("executing RegisterAllFHIRResourceRoutes")
	for _, i := range registry {
		registerFHIRResourceRoutes(r, log, rndr, i)
	}
}

// RegisterHealthCheckRoutes ...
func RegisterHealthCheckRoutes(r *mux.Router, log *logging.Logger) {
	log.Debug("executing RegisterHealthCheckRoutes")
	hc := healthcheck.NewHandler()
	r.HandleFunc("/live", hc.LiveEndpoint)
	r.HandleFunc("/ready", hc.ReadyEndpoint)
}

func registerFHIRResourceRoutes(r *mux.Router, log *logging.Logger, rndr *render.Render, i interface{}) {
	dLog := log.WithField("function", "registerFHIRResourceRoutes")
	dLog.Debug("executing")

	seg := utils.GetBaseTypeName(i)
	dLog = dLog.WithField("resource", seg)

	typePrefix := "/" + seg
	instancePrefix := typePrefix + "/{resourceID}"

	if t, ok := i.(resources.CreateableResource); ok {
		dLog.Debug("registering create method")
		r.Handle(typePrefix, t.Create(log, rndr)).Methods("POST")
	}

	if t, ok := i.(resources.SearchableResource); ok {
		dLog.Debug("registering search method")
		r.Handle(typePrefix, t.Search(log, rndr)).Methods("GET")
	}

	if t, ok := i.(resources.UpdateableResource); ok {
		dLog.Debug("registering update method")
		r.Handle(instancePrefix, t.Update(log, rndr)).Methods("PUT")
	}

	if t, ok := i.(resources.DeleteableResource); ok {
		dLog.Debug("registering delete method")
		r.Handle(instancePrefix, t.Delete(log, rndr)).Methods("DELETE")
	}

	if t, ok := i.(resources.ReadableResource); ok {
		dLog.Debug("registering read method")
		r.Handle(instancePrefix, t.Read(log, rndr)).Methods("GET")
	}

	if t, ok := i.(resources.PatchableResource); ok {
		dLog.Debug("registering patch method")
		r.Handle(instancePrefix, t.Patch(log, rndr)).Methods("PATCH")
	}

	if t, ok := i.(resources.VersionReadableResource); ok {
		dLog.Debug("registering version read method")
		r.Handle(instancePrefix+"/_history/{versionID}", t.VersionRead(log, rndr)).Methods("GET")
	}

	if t, ok := i.(resources.ValidateableResource); ok {
		dLog.Debug("registering validate method")
		r.Handle(typePrefix+"/$validate", t.Validate(log, rndr)).Methods("POST")
	}
}
