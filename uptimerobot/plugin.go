package uptimerobot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Plugin creates this (uptimerobot) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-uptimerobot",
		DefaultTransform: transform.FromCamel(),
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"not found"}),
			},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "user_id",
				Hydrate: getAccountUserId,
			},
		},
		TableMap: map[string]*plugin.Table{
			"uptimerobot_account":            tableUptimeRobotAccount(ctx),
			"uptimerobot_alert_contact":      tableUptimeRobotAlertContact(ctx),
			"uptimerobot_maintenance_window": tableUptimeRobotMaintenanceWindow(ctx),
			"uptimerobot_monitor":            tableUptimeRobotMonitor(ctx),
		},
	}
	return p
}
