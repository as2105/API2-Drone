package config

import (
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/ethereum/go-ethereum/common"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/oliveagle/jsonpath"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ObjectCollectionContract ...
type ObjectCollectionContract struct {
	Name    string
	Address common.Address
	Indexes []*ObjectIndex
}

// ObjectIndex ...
type ObjectIndex struct {
	Name     string
	Address  common.Address
	JSONPath *jsonpath.Compiled
}

// Config ...
type Config struct {
	Address                   string   `mapstructure:"address"`
	CORSAllowCredentials      bool     `mapstructure:"cors_allow_credentials"`
	CORSAllowedHeaders        []string `mapstructure:"cors_allowed_headers"`
	CORSAllowedMethods        []string `mapstructure:"cors_allowed_methods"`
	CORSAllowedOrigins        []string `mapstructure:"cors_allowed_origins"`
	CORSExposedHeaders        []string `mapstructure:"cors_exposed_headers"`
	CORSMaxAge                int      `mapstructure:"cors_max_age"`
	DatabaseConnectionString  string   `mapstructure:"db_conn_str"`
	DatabaseType              string   `mapstructure:"db_type"`
	DevMode                   bool     `mapstructure:"dev_mode"`
	GasLimit                  uint64   `mapstructure:"gas_limit"`
	GasPrice                  int64    `mapstructure:"gas_price"`
	LogFormat                 string   `mapstructure:"log_format"`
	LogLevel                  string   `mapstructure:"log_level"`
	PrivateKey                string   `mapstructure:"private_key"`
	Profile                   bool     `mapstructure:"profile"`
	RPCURL                    string   `mapstructure:"rpc_url"`
	TransactionsChannelBuffer uint     `mapstructure:"txns_buffer"`
	Pprof                    bool              `mapstructure:"pprof"`

	OrganizationContract      common.Address
	ObjectCollectionContracts map[string]*ObjectCollectionContract
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

	// special contracts configuration setup
	c.OrganizationContract = common.HexToAddress(viper.GetString("contracts.organization.address"))
	contractsConfig, err := parseContractsConfiguration(viper.GetStringMap("contracts"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse contracts configuration")
	}
	c.ObjectCollectionContracts = contractsConfig

	// set up logger
	logging.SetLevel(log, c.LogLevel)
	logging.SetFormatter(log, c.LogFormat)

	log.WithFields(viper.AllSettings()).Debug("config")
	log.Debug("Config object: ", c)

	return c, nil
}

func parseContractsConfiguration(in map[string]interface{}) (map[string]*ObjectCollectionContract, error) {
	newMap := map[string]*ObjectCollectionContract{}
	rawContractsConfig, ok := in["collections"]
	if !ok {
		return newMap, errors.New("failed to find contracts configuration")
	}
	rawCollArr := rawContractsConfig.([]interface{})
	for _, rawColl := range rawCollArr {
		collData := rawColl.(map[interface{}]interface{})
		cName := collData["name"].(string)
		cAddr := common.HexToAddress(collData["address"].(string))
		idxColl := []*ObjectIndex{}
		if rawIdxs, ok := collData["indexes"].([]interface{}); ok {
			for _, rawIdx := range rawIdxs {
				idxData := rawIdx.(map[interface{}]interface{})
				idxName := idxData["name"].(string)
				idxAddr := common.HexToAddress(idxData["address"].(string))
				idxPath := idxData["path"].(string)
				idxPathCompiled, err := jsonpath.Compile(idxPath)
				if err != nil {
					return newMap, errors.Wrapf(err, "unable to compile JSON path %q", idxPath)
				}
				newIdx := ObjectIndex{
					Name:     idxName,
					Address:  idxAddr,
					JSONPath: idxPathCompiled,
				}
				idxColl = append(idxColl, &newIdx)
			}
		}
		newMap[cName] = &ObjectCollectionContract{
			Name:    cName,
			Address: cAddr,
			Indexes: idxColl,
		}
	}
	return newMap, nil
}

// BindFlags ...
func BindFlags(c *cobra.Command) {
	viper.BindPFlags(c.Flags())
}
