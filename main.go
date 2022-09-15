package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-uptimerobot/uptimerobot"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: uptimerobot.Plugin})
}
