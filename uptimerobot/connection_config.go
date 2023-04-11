package uptimerobot

import (
	"context"
	"errors"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
	"github.com/turbot/uptimerobotapi"
)

type uptimerobotConfig struct {
	APIKey *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
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

	apiKey := os.Getenv("UPTIMEROBOT_API_KEY")
	if uptimerobotConfig.APIKey != nil {
		apiKey = *uptimerobotConfig.APIKey
	}
	if apiKey == "" {
		return nil, errors.New("'api_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	client := uptimerobotapi.NewClient(apiKey)

	return client, nil
}
