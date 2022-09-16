package uptimerobot

import (
	"context"

	"github.com/bigdatasourav/uptimerobotapi"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableUptimeRobotMaintenanceWindow(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_maintenance_window",
		Description: "UptimeRobot Maintenance Window.",
		List: &plugin.ListConfig{
			Hydrate: listMaintenanceWindows,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "type",
					Require: plugin.Optional,
				},
				{
					Name:    "duration",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique id of the maintenance window."},
			{Name: "user", Type: proto.ColumnType_INT, Description: "User of the maintenance window."},
			{Name: "type", Type: proto.ColumnType_INT, Description: "Type of the maintenance window."},
			{Name: "friendly_name", Type: proto.ColumnType_STRING, Description: "Friendly name of the maintenance window."},
			{Name: "start_time", Type: proto.ColumnType_STRING, Description: "Start time of the maintenance window."},
			{Name: "duration", Type: proto.ColumnType_INT, Description: "Duration of the maintenance window."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Value of the maintenance window."},
			{Name: "status", Type: proto.ColumnType_INT, Description: "Status of the maintenance window."},
		},
	}
}

func listMaintenanceWindows(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listMaintenanceWindows", "conection_error", err)
		return nil, err
	}
	var params = uptimerobotapi.GetMWindowParams{}

	windowType := d.KeyColumnQuals["type"].GetStringValue()
	if windowType != "" {
		params.Type = windowType
	}

	duration := d.KeyColumnQuals["duration"].GetStringValue()
	if duration != "" {
		params.Duration = duration
	}

	mw, err := conn.MWindow.GetMWindows(params)
	if err != nil {
		plugin.Logger(ctx).Error("listMaintenanceWindows", "api_error", err)
		return nil, err
	}

	for _, element := range mw.MWindows {
		d.StreamListItem(ctx, element)
	}
	return nil, nil
}
