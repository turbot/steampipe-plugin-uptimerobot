package uptimerobot

import (
	"context"
	"errors"
	"os"

	uptimerobot "github.com/bitfield/uptimerobot/pkg"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
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

func connect(_ context.Context, d *plugin.QueryData) (uptimerobot.Client, error) {
	uptimerobotConfig := GetConfig(d.Connection)

	apiKey := os.Getenv("UPTIME_ROBOT_API_KEY")
	if uptimerobotConfig.APIKey != nil {
		apiKey = *uptimerobotConfig.APIKey
	}
	if apiKey == "" {
		test := uptimerobot.Client{}
		return test, errors.New("'api_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	client := uptimerobot.New(apiKey)

	return client, nil
}
