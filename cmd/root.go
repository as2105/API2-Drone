// Copyright © 2018 Optum

package cmd

import (
	"github.com/SynapticHealthAlliance/fhir-api/config"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{Use: "fhir-api"}
	log     = logging.NewLogger()
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("unable to set up root command")
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.fhir-api.yaml)")

	rootCmd.PersistentFlags().StringP("log_level", "l", "info", "logging level (trace, debug, info, warn, error, fatal)")
	rootCmd.PersistentFlags().StringP("log_format", "f", "default", "logging format (default or json)")

	rootCmd.PersistentFlags().String("db_type", "sqlite3", "database dialect to use (sqlite3 or postgresql)")
	rootCmd.PersistentFlags().String("db_conn_str", "data.sqlite", "database dialect-specific connection string")

	rootCmd.PersistentFlags().Int64P("gas_price", "g", 0, "gas price to use for transactions")
	rootCmd.PersistentFlags().Uint64P("gas_limit", "L", 0, "gas limit to use for transactions")
	rootCmd.PersistentFlags().StringP("private_key", "k", "", "private key for signing transactions")
	rootCmd.PersistentFlags().StringP("rpc_url", "u", "http://localhost:8545", "ethereum node RPC URL")
	rootCmd.PersistentFlags().StringP("org_name", "n", "", "name of your Organization, to be attached to Organization contract")
	rootCmd.PersistentFlags().StringP("org_contract_addr", "O", "", "address of this Organization's smart contract")

	rootCmd.PersistentFlags().Bool("dev_mode", false, "development mode")
	rootCmd.PersistentFlags().Uint("txns_buffer", 2048, "")

	initServe()
}

func initConfig() {
	config.GlobalSettings.ConfigFile = cfgFile
}
