// Copyright © 2018 Optum

package cmd

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/config"
	"github.com/SynapticHealthAlliance/fhir-api/handlers"
	"github.com/SynapticHealthAlliance/fhir-api/handlers/resources"
	"github.com/SynapticHealthAlliance/fhir-api/internal/metadata"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/SynapticHealthAlliance/fhir-api/static"
	"github.com/SynapticHealthAlliance/fhir-api/storage/database"
	"github.com/SynapticHealthAlliance/fhir-api/storage/ethereum"
	"github.com/SynapticHealthAlliance/fhir-api/utils"
	"github.com/gorilla/mux"
	nLog "github.com/meatballhat/negroni-logrus"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
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
	serveCmd.Flags().StringP("address", "a", "localhost:8080", "interface and port on which the server will run")
	serveCmd.Flags().StringArray("cors_allowed_origins", []string{"*"}, "")
	serveCmd.Flags().StringArray("cors_allowed_methods", []string{"GET", "POST", "PUT", "PATCH", "DELETE"}, "")
	serveCmd.Flags().StringArray("cors_allowed_headers", []string{}, "")
	serveCmd.Flags().StringArray("cors_exposed_headers", []string{}, "")
	serveCmd.Flags().Bool("cors_allow_credentials", false, "")
	serveCmd.Flags().Int("cors_max_age", 0, "")
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
			resources.NewResourceRegistry,
			newRenderer,
			newCORSMiddleware,
			ethereum.NewConnection,
			ethereum.NewTransactOpts,
			ethereum.NewTransactionsChannel,
			ethereum.NewTransactionsListener,
			database.NewConnection,
			static.NewStaticFilesBox,
		),
		fx.Logger(logging.NewLogger()),
		fx.Invoke(
			configureRouter,
			ethereum.StartTransactionsListener,
		),
	)

	app.Run()
}

func newCORSMiddleware(config *config.Config) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   config.CORSAllowedOrigins,
		AllowedMethods:   config.CORSAllowedMethods,
		AllowedHeaders:   config.CORSExposedHeaders,
		ExposedHeaders:   config.CORSExposedHeaders,
		AllowCredentials: config.CORSAllowCredentials,
		MaxAge:           config.CORSMaxAge,
		Debug:            config.DevMode,
	})
}

func newRenderer(config *config.Config) *render.Render {
	r := render.New(render.Options{
		IndentJSON:      true,
		IndentXML:       true,
		IsDevelopment:   config.DevMode,
		JSONContentType: fmt.Sprintf("application/fhir+json; fhirVersion=%s", models.FHIRVersion),
	})
	return r
}

// NewRouter ...
func NewRouter(lc fx.Lifecycle, log *logging.Logger, config *config.Config, corsMW *cors.Cors) *mux.Router {
	log.Debug("executing NewRouter")

	r := mux.NewRouter()
	n := negroni.New()

	// enable CORS
	n.Use(corsMW)

	// recover on panics and throw a 500
	recovery := negroni.NewRecovery()
	recovery.PrintStack = config.DevMode
	n.Use(recovery)

	// request logging
	n.Use(nLog.NewMiddlewareFromLogger(log, "web"))

	// response gzip
	n.Use(gzip.Gzip(gzip.DefaultCompression))

	n.UseHandler(r)

	server := &http.Server{
		Handler:      n,
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

func configureRouter(r *mux.Router, log *logging.Logger, cConfig *resources.CapabilityConfig, rndr *render.Render, registry resources.ResourceRegistry) {
	log.Debug("executing configureRouter")

	fhirRouter := r.PathPrefix("/fhir").Subrouter()
	handlers.RegisterFHIRCapabilityStatementRoutes(fhirRouter, log, cConfig, rndr)
	handlers.RegisterAllFHIRResourceRoutes(fhirRouter, log, rndr, registry)

	handlers.RegisterHealthCheckRoutes(r, log)

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
