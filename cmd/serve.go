// Copyright © 2018 Optum

package cmd

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/config"
	"github.com/SynapticHealthAlliance/fhir-api/handlers"
	"github.com/SynapticHealthAlliance/fhir-api/handlers/resources"
	"github.com/SynapticHealthAlliance/fhir-api/internal/metadata"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/utils"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// serveCmd represents the serve command
var serveCmd *cobra.Command

func initServe() {
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		Run:   serveRun,
	}
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("address", "a", "localhost:8080", "Interface and port on which the server will run")
}

func serveRun(cmd *cobra.Command, args []string) {
	log.WithFields(utils.StructToMap(metadata.Metadata)).Info("build information")

	config.BindFlags(serveCmd)

	app := fx.New(
		fx.Provide(
			logging.NewLogger,
			config.NewConfig,
			NewRouter,
			resources.NewCapabilityConfig,
		),
		fx.Logger(logging.NewLogger()),
		fx.Invoke(
			configureRouter,
		),
	)

	app.Run()
}

// NewRouter ...
func NewRouter(lc fx.Lifecycle, log *logging.Logger, config *config.Config) *mux.Router {
	log.Debug("executing NewRouter")

	r := mux.NewRouter()

	server := &http.Server{
		Handler:      r,
		Addr:         config.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Info("starting server")
			go func() {
				if err := server.ListenAndServe(); err != nil {
					log.WithError(err).Error("server failed")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("stopping server")
			return server.Shutdown(ctx)
		},
	})

	return r
}

func configureRouter(r *mux.Router, log *logging.Logger, cConfig *resources.CapabilityConfig) {
	log.Debug("executing configureRouter")
	fhirRouter := r.PathPrefix("/fhir").Subrouter()
	handlers.RegisterFHIRCapabilityStatementRoutes(fhirRouter, log, cConfig)
	handlers.RegisterAllFHIRResourceRoutes(fhirRouter, log)
	debugRouter(r, log)
}

func debugRouter(r *mux.Router, log *logging.Logger) {
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		f := logging.Fields{}
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			f["route"] = pathTemplate
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			f["regexp"] = pathRegexp
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			f["queriesTemplates"] = strings.Join(queriesTemplates, ",")
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			f["queriesRegexps"] = strings.Join(queriesRegexps, ",")
		}
		methods, err := route.GetMethods()
		if err == nil {
			f["methods"] = strings.Join(methods, ",")
		}
		log.WithFields(f).Debug("route registered")
		return nil
	})
}
