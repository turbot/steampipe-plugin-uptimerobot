package uptimerobot

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type uptimerobotConfig struct {
	Token *string `cty:"apiKey"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"apiKey": {
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
