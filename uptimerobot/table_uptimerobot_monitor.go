package uptimerobot

import (
	"context"

	"github.com/bigdatasourav/uptimerobotapi"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableUptimerobotMonitor(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_monitor",
		Description: "UptimeRobot Retrieve Websites Monitored Current Account.",
		List: &plugin.ListConfig{
			Hydrate: listMonitors,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getMonitor,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique account id for monitors."},
			{Name: "friendly_name", Type: proto.ColumnType_STRING, Description: "The friendly name for monitors."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "The url/ip of the monitored website."},
			{Name: "type", Type: proto.ColumnType_INT, Description: "Type of the website monitored."},
			{Name: "sub_type", Type: proto.ColumnType_STRING, Description: "Subtype of the website monitored."},
			{Name: "keyword_type", Type: proto.ColumnType_STRING, Description: "The keyword of the monitors."},
			{Name: "keyword_case_type", Type: proto.ColumnType_INT, Description: "The keyword case of the monitors."},
			{Name: "port", Type: proto.ColumnType_STRING, Description: "The port of the website monitored."},
			{Name: "keyword_value", Type: proto.ColumnType_STRING, Description: "The keyword value of the website."},
			{Name: "http_username", Type: proto.ColumnType_STRING, Description: "The http-username for password-protected web pages."},
			{Name: "http_password", Type: proto.ColumnType_STRING, Description: "The http-password for password-protected web pages."},
			{Name: "interval", Type: proto.ColumnType_INT, Description: "The interval for the monitoring check ."},
			{Name: "timeout", Type: proto.ColumnType_INT, Description: "The timeout for the monitoring check ."},
			{Name: "create_datetime", Type: proto.ColumnType_TIMESTAMP, Description: "The creation time for the monitor.", Transform: transform.FromField("CreateDatetime.Time")},
			{Name: "monitor_group", Type: proto.ColumnType_INT, Description: "The groups for monitors."},
			{Name: "is_group_main", Type: proto.ColumnType_INT, Description: "The main groups for monitors."},
			{Name: "alert_contacts", Type: proto.ColumnType_JSON, Description: "The alert contacts of the monitors."},
			{Name: "logs", Type: proto.ColumnType_JSON, Description: "The logs of the monitors."},
			{Name: "status", Type: proto.ColumnType_INT, Description: "The status of the monitored website."},
			{Name: "ssl", Type: proto.ColumnType_JSON, Description: "The ssl of the monitors."},
		},
	}
}

func listMonitors(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listMonitor", "connection_error", err)
		return nil, err
	}

	input := uptimerobotapi.GetMonitorsParams{Limit: types.Int(50)}
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < int64(*input.Limit) {
			if *limit < 1 {
				input.Limit = types.Int(1)
			} else {
				input.Limit = types.Int(int(*limit))
			}
		}
	}

	monitors, err := conn.Monitor.GetMonitors(input)
	if err != nil {
		plugin.Logger(ctx).Error("listMonitor", "api_error", err)
		return nil, err
	}
	for _, monitor := range monitors.Monitors {
		d.StreamListItem(ctx, monitor)
	}
	return nil, nil
}

func getMonitor(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getMonitor", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()
	if id == "" {
		return nil, nil
	}
	result, err := conn.Monitor.GetMonitors(uptimerobotapi.GetMonitorsParams{Monitors: &id})
	if err != nil {
		plugin.Logger(ctx).Error("getMonitor", "api_error", err)
		return nil, err
	}
	if len(result.Monitors) > 0 {
		return result.Monitors[0], nil
	}

	return nil, nil
}
