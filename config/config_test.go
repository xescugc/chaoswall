package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xescugc/chaoswall/config"
)

func TestNew(t *testing.T) {
	tests := []struct {
		Name   string
		Config config.Config
		Error  string
	}{
		{
			Name: "Success",
			Config: config.Config{
				DBHost:     "dbhost",
				DBPort:     8080,
				DBUser:     "dbuser",
				DBPassword: "dbpass",
				DBName:     "dbname",
				Port:       3030,
			},
		},
		{
			Name:   "ErrorDBHost",
			Config: config.Config{},
			Error:  "DBHost is required",
		},
		{
			Name: "ErrorDBPort",
			Config: config.Config{
				DBHost: "dbhost",
			},
			Error: "DBPort is required",
		},
		{
			Name: "ErrorDBUser",
			Config: config.Config{
				DBHost: "dbhost",
				DBPort: 8080,
			},
			Error: "DBUser is required",
		},
		{
			Name: "ErrorDBPassword",
			Config: config.Config{
				DBHost: "dbhost",
				DBPort: 8080,
				DBUser: "dbuser",
			},
			Error: "DBPassword is required",
		},
		{
			Name: "ErrorDBName",
			Config: config.Config{
				DBHost:     "dbhost",
				DBPort:     8080,
				DBUser:     "dbuser",
				DBPassword: "dbpass",
			},
			Error: "DBName is required",
		},
		{
			Name: "ErrorPort",
			Config: config.Config{
				DBHost:     "dbhost",
				DBPort:     8080,
				DBUser:     "dbuser",
				DBPassword: "dbpass",
				DBName:     "dbname",
			},
			Error: "Port is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := tt.Config.Validate()
			if tt.Error == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.Error)
			}
		})
	}
}
