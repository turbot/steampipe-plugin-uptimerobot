package main

import (
	"github.com/turbot/steampipe-plugin-uptimerobot/uptimerobot"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: uptimerobot.Plugin})
}
