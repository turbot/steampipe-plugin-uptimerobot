package uptimerobot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableUptimerobotMonitor(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_monitor",
		Description: "Retrieve monitoring websites in your current account.",
		List: &plugin.ListConfig{
			Hydrate: listMonitors,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getMonitors,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "The unique account id for monitors.", Transform: transform.FromField("ID")},
			{Name: "friendly_name", Type: proto.ColumnType_STRING, Description: "The friendly name for monitors."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "The url/ip of the monitored website.", Transform: transform.FromField("URL")},
			{Name: "type", Type: proto.ColumnType_INT, Description: "Type of the website monitored."},
			{Name: "sub_type", Type: proto.ColumnType_INT, Description: "Subtype of the website monitored."},
			{Name: "keyword_type", Type: proto.ColumnType_INT, Description: "The keyword of the monitors."},
			{Name: "port", Type: proto.ColumnType_INT, Description: "The port of the website monitored."},
			{Name: "keyword_value", Type: proto.ColumnType_STRING, Description: "The keyword value of the website."},
			{Name: "alert_contacts", Type: proto.ColumnType_STRING, Description: "The alert contacts of the monitors."},
			{Name: "status", Type: proto.ColumnType_INT, Description: "The status of the monitored website."},
		},
	}
}

func listMonitors(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listMonitor", "connection_error", err)
		return nil, err
	}

	monitors, err := conn.AllMonitors()
	if err != nil {
		plugin.Logger(ctx).Error("listMonitor", "api_error", err)
		return nil, err
	}

	for _, monitor := range monitors {
		d.StreamListItem(ctx, monitor)
	}

	return nil, nil
}
func getMonitors(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getMonitor", "connection_error", err)
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetInt64Value()

	result, err := conn.GetMonitor(id)
	if err != nil {
		plugin.Logger(ctx).Error("getMonitor", "api_error", err)
		return nil, err
	}
	return result, nil
}
