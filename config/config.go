package config

import (
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	Address                  string            `mapstructure:"address"`
	ContractAddresses        map[string]string `mapstructure:"contract_addresses"`
	CORSAllowCredentials     bool              `mapstructure:"cors_allow_credentials"`
	CORSAllowedHeaders       []string          `mapstructure:"cors_allowed_headers"`
	CORSAllowedMethods       []string          `mapstructure:"cors_allowed_methods"`
	CORSAllowedOrigins       []string          `mapstructure:"cors_allowed_origins"`
	CORSExposedHeaders       []string          `mapstructure:"cors_exposed_headers"`
	CORSMaxAge               int               `mapstructure:"cors_max_age"`
	DatabaseConnectionString string            `mapstructure:"db_conn_str"`
	DatabaseType             string            `mapstructure:"db_type"`
	DevMode                  bool              `mapstructure:"dev_mode"`
	GasPrice                 int64             `mapstructure:"gas_price"`
	LogFormat                string            `mapstructure:"log_format"`
	LogLevel                 string            `mapstructure:"log_level"`
	PrivateKey               string            `mapstructure:"private_key"`
	Profile                  bool              `mapstructure:"profile"`
	RPCURL                   string            `mapstructure:"rpcurl"`
}

// GlobalSettings holds settings configured by global flags, that may
var GlobalSettings = struct {
	ConfigFile string
}{}

// NewConfig ...
func NewConfig(log *logging.Logger) (*Config, error) {
	c := &Config{}

	if GlobalSettings.ConfigFile != "" {
		viper.SetConfigFile(GlobalSettings.ConfigFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return nil, errors.Wrap(err, "unable to determine home directory")
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".fhir-api")
	}

	viper.SetEnvPrefix("FHIRAPI")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "unable to read config file")
	}
	log.Info("Using config file:", viper.ConfigFileUsed())

	if err := viper.Unmarshal(c); err != nil {
		return nil, errors.Wrap(err, "unable to read configuration")
	}

	// set up logger
	logging.SetLevel(log, c.LogLevel)
	logging.SetFormatter(log, c.LogFormat)

	log.WithFields(viper.AllSettings()).Debug("config")

	return c, nil
}

// BindFlags ...
func BindFlags(c *cobra.Command) {
	viper.BindPFlags(c.Flags())
}
