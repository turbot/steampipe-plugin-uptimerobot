package uptimerobot

import (
	"context"
	"errors"
	"os"

	uptimerobotapi "github.com/bigdatasourav/uptimerobotapi"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type uptimerobotConfig struct {
	APIToken *string `cty:"api_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &uptimerobotConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) uptimerobotConfig {
	if connection == nil || connection.Config == nil {
		return uptimerobotConfig{}
	}
	config, _ := connection.Config.(uptimerobotConfig)
	return config
}

func connect(_ context.Context, d *plugin.QueryData) (*uptimerobotapi.Client, error) {
	uptimerobotConfig := GetConfig(d.Connection)

	apiToken := os.Getenv("UPTIMEROBOT_API_TOKEN")
	if uptimerobotConfig.APIToken != nil {
		apiToken = *uptimerobotConfig.APIToken
	}
	if apiToken == "" {
		return nil, errors.New("'api_token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	client := uptimerobotapi.NewClient(apiToken)

	return client, nil
}
