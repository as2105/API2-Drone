// Copyright © 2018 Optum
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: serveRun,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", 8000, "TCP port to host traffic on")
	viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
}

func serveRun(cmd *cobra.Command, args []string) {
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, os.Kill)

	app := fx.New(
		fx.Provide(
			NewHandler,
			NewRouter,
			logrus.StandardLogger,
		),
		fx.Logger(logrus.StandardLogger()),
		fx.Invoke(
			RegisterFHIRCapabilityStatement,
			RegisterFHIRPractitionerAPI,
			RegisterFHIRPractitionerRoleAPI,
			RegisterFHIRLocationAPI,
		),
	)

	app.Run()
	<-app.Done()
}

// NewHandler constructs a simple HTTP handler. Since it returns an
// http.Handler, Fx will treat NewHandler as the constructor for the
// http.Handler type.
//
// Like many Go functions, NewHandler also returns an error. If the error is
// non-nil, Go convention tells the caller to assume that NewHandler failed
// and the other returned values aren't safe to use. Fx understands this
// idiom, and assumes that any function whose last return value is an error
// follows this convention.
//
// Unlike NewLogger, NewHandler has formal parameters. Fx will interpret these
// parameters as dependencies: in order to construct an HTTP handler,
// NewHandler needs a logger. If the application has access to a *log.Logger
// constructor (like NewLogger above), it will use that constructor or its
// cached output and supply a logger to NewHandler. If the application doesn't
// know how to construct a logger and needs an HTTP handler, it will fail to
// start.
//
// Functions may also return multiple objects. For example, we could combine
// NewHandler and NewLogger into a single function:
//
//   func NewHandlerAndLogger() (*log.Logger, http.Handler, error)
//
// Fx also understands this idiom, and would treat NewHandlerAndLogger as the
// constructor for both the *log.Logger and http.Handler types. Just like
// constructors for a single type, NewHandlerAndLogger would be called at most
// once, and both the handler and the logger would be cached and reused as
// necessary.
func NewHandler(logger *logrus.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a request.")
	}), nil
}

// NewRouter constructs an HTTP mux. Like NewHandler, it depends on *log.Logger.
// However, it also depends on the Fx-specific Lifecycle interface.
//
// A Lifecycle is available in every Fx application. It lets objects hook into
// the application's start and stop phases. In a non-Fx application, the main
// function often includes blocks like this:
//
//   srv, err := NewServer() // some long-running network server
//   if err != nil {
//     log.Fatalf("failed to construct server: %v", err)
//   }
//   // Construct other objects as necessary.
//   go srv.Start()
//   defer srv.Stop()
//
// In this example, the programmer explicitly constructs a bunch of objects,
// crashing the program if any of the constructors encounter unrecoverable
// errors. Once all the objects are constructed, we start any background
// goroutines and defer cleanup functions.
//
// Fx removes the manual object construction with dependency injection. It
// replaces the inline goroutine spawning and deferred cleanups with the
// Lifecycle type.
//
// Here, NewMux makes an HTTP mux available to other functions. Since
// constructors are called lazily, we know that NewMux won't be called unless
// some other function wants to register a handler. This makes it easy to use
// Fx's Lifecycle to start an HTTP server only if we have handlers registered.
func NewRouter(lc fx.Lifecycle, logger *logrus.Logger) *mux.Router {
	logger.Print("Executing NewRouter.")

	r := mux.NewRouter()
	server := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", viper.GetInt("port")),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// If NewMux is called, we know that another function is using the mux. In
	// that case, we'll use the Lifecycle type to register a Hook that starts
	// and stops our HTTP server.
	//
	// Hooks are executed in dependency order. At startup, NewLogger's hooks run
	// before NewMux's. On shutdown, the order is reversed.
	//
	// Returning an error from OnStart hooks interrupts application startup. Fx
	// immediately runs the OnStop portions of any successfully-executed OnStart
	// hooks (so that types which started cleanly can also shut down cleanly),
	// then exits.
	//
	// Returning an error from OnStop hooks logs a warning, but Fx continues to
	// run the remaining hooks.
	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 30 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return r
}

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

// RegisterFHIRPractitionerAPI mounts our HTTP handler on the mux.
func RegisterFHIRPractitionerAPI(r *mux.Router, logger *logrus.Logger, h http.Handler) {
	s := r.PathPrefix("/Practitioner").Subrouter()
	s.Handle("/", newCreatePractitionerHandler(logger)).Methods("POST")
	s.Handle("/{practitionerId}", newReadPractitionerHandler(logger)).Methods("GET")
	s.Handle("/{practitionerId}", newUpdatePractitionerHandler(logger)).Methods("PUT")
}

func newCreatePractitionerHandler(logger *logrus.Logger) http.Handler {
	logger.Print("Executing newCreatePractitionerHandler.")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var u struct{}
		if r.Body == nil {
			http.Error(rw, "Please send a request body", http.StatusBadRequest)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		// headers := rw.Header()
		// headers.Set("Location", fmt.Sprintf("[base]/[type]/[id]/_history/[vid]"))
		// headers.Set("ETag", fmt.Sprintf("")) // versionId (if versioning is supported)
		// headers.Set("Last-Modified", time.Now().Format(time.RFC1123Z))

		// 400 http.StatusBadRequest - resource could not be parsed or failed basic FHIR validation rules
		// 404 http.StatusNotFound - resource type not supported, or not a FHIR end-point
		// 422 http.StatusUnprocessableEntity - the proposed resource violated applicable FHIR profiles or server business rules. This should be accompanied by an OperationOutcome resource providing additional detail

		if ifNoneExist, ccreate := r.Header["If-None-Exist"]; ccreate {
			http.Error(rw, fmt.Sprintf("Unsupported Header: \"If-None-Exist\": %v", ifNoneExist), http.StatusUnprocessableEntity)
		}
		// conditional create:
		// If-None-Exist: identifier=http://my-lab-system|123
		//                           http://hl7.org/fhir/sid/us-npi
		//                           http://hl7.org/fhir/sid/us-ssn
		//                           urn:ietf:rfc:3986
		// No matches: The server processes the create as above
		// One Match: The server ignore the post and returns 200 OK
		// Multiple matches: The server returns a 412 Precondition Failed error indicating the client's criteria were not selective enough. return an OperationOutcome
	})
}

func newReadPractitionerHandler(logger *logrus.Logger) http.Handler {
	logger.Print("Executing newReadPractitionerHandler.")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		v := map[string]string{
			"message": "hello world",
		}
		json.NewEncoder(rw).Encode(v)
	})
}

func newUpdatePractitionerHandler(logger *logrus.Logger) http.Handler {
	logger.Print("Executing newUpdatePractitionerHandler.")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var u struct{}
		if r.Body == nil {
			http.Error(rw, "Please send a request body", http.StatusBadRequest)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
	})
}

// RegisterFHIRPractitionerRoleAPI mounts our HTTP handler on the mux.
func RegisterFHIRPractitionerRoleAPI(r *mux.Router, logger *logrus.Logger, h http.Handler) {
	s := r.PathPrefix("/PractitionerRole").Subrouter()
	s.Handle("/", newCreatePractitionerRoleHandler(logger)).Methods("POST")
	s.Handle("/{practitionerRoleId}", newReadPractitionerRoleHandler(logger)).Methods("GET")
	s.Handle("/{practitionerRoleId}", newUpdatePractitionerRoleHandler(logger)).Methods("PUT")
	s.Handle("/{practitionerRoleId}", newDeletePractitionerRoleHandler(logger)).Methods("DELETE")
}

func newCreatePractitionerRoleHandler(logger *logrus.Logger) http.Handler {
	logger.Print("Executing newCreatePractitionerRoleHandler.")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var u struct{}
		if r.Body == nil {
			http.Error(rw, "Please send a request body", http.StatusBadRequest)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
	})
}

func newReadPractitionerRoleHandler(logger *logrus.Logger) http.Handler {
	logger.Print("Executing newReadPractitionerRoleHandler.")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		v := map[string]string{
			"message": "hello world",
		}
		json.NewEncoder(rw).Encode(v)
	})
}

func newUpdatePractitionerRoleHandler(logger *logrus.Logger) http.Handler {
	logger.Print("Executing newUpdatePractitionerRoleHandler.")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var u struct{}
		if r.Body == nil {
			http.Error(rw, "Please send a request body", http.StatusBadRequest)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
	})
}

func newDeletePractitionerRoleHandler(logger *logrus.Logger) http.Handler {
	logger.Print("Executing newDeletePractitionerRoleHandler.")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		v := map[string]string{
			"message": "hello world",
		}
		json.NewEncoder(rw).Encode(v)
	})
}

// RegisterFHIRLocationAPI mounts our HTTP handler on the mux.
func RegisterFHIRLocationAPI(r *mux.Router, logger *logrus.Logger, h http.Handler) {
	s := r.PathPrefix("/Location").Subrouter()
	s.Handle("/", newCreateLocationHandler(logger)).Methods("POST")
	s.Handle("/{locationId}", newReadLocationHandler(logger)).Methods("GET")
	s.Handle("/{locationId}", newUpdateLocationHandler(logger)).Methods("PUT")
}

func newCreateLocationHandler(logger *logrus.Logger) http.Handler {
	logger.Print("Executing newCreateLocationHandler.")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var u struct{}
		if r.Body == nil {
			http.Error(rw, "Please send a request body", http.StatusBadRequest)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
	})
}

func newReadLocationHandler(logger *logrus.Logger) http.Handler {
	logger.Print("Executing newReadLocationHandler.")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		v := map[string]string{
			"message": "hello world",
		}
		json.NewEncoder(rw).Encode(v)
	})
}

func newUpdateLocationHandler(logger *logrus.Logger) http.Handler {
	logger.Print("Executing newUpdateLocationHandler.")
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var u struct{}
		if r.Body == nil {
			http.Error(rw, "Please send a request body", http.StatusBadRequest)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
	})
}
