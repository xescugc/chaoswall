package config

import (
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/xerrors"
)

const (
	DefaultPort = 9090
)

type Config struct {
	// Port is the port to open to the public
	Port int `mapstructure:"port"`

	DBHost     string `mapstructure:"db-host"`
	DBPort     int    `mapstructure:"db-port"`
	DBUser     string `mapstructure:"db-user"`
	DBPassword string `mapstructure:"db-password"`
	DBName     string `mapstructure:"db-name"`
}

// New returns a new Config from the viper.Viper, the ENV variables
// are readed by using the convertion of "_" and all caps
func New(v *viper.Viper) (*Config, error) {
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	if v.GetString("config") != "" {
		v.SetConfigFile(v.GetString("config"))
		err := v.ReadInConfig()
		if err != nil {
			return nil, err
		}
	}

	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Validate validates the Config
func (c *Config) Validate() error {
	if c.DBHost == "" {
		return xerrors.New("DBHost is required")
	} else if c.DBPort == 0 {
		return xerrors.New("DBPort is required")
	} else if c.DBUser == "" {
		return xerrors.New("DBUser is required")
	} else if c.DBPassword == "" {
		return xerrors.New("DBPassword is required")
	} else if c.DBName == "" {
		return xerrors.New("DBName is required")
	} else if c.Port == 0 {
		return xerrors.New("Port is required")
	}

	return nil
}
