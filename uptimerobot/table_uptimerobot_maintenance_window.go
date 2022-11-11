package uptimerobot

import (
	"context"
	"strconv"

	"github.com/bigdatasourav/uptimerobotapi"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableUptimeRobotMaintenanceWindow(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_maintenance_window",
		Description: "Uptime Robot has a new feature (for the Pro Plan) to handle one-time or regular downtimes nicely.",
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
			{Name: "id", Type: proto.ColumnType_INT, Description: "The ID of the maintenance window."},
			{Name: "friendly_name", Type: proto.ColumnType_STRING, Description: "Friendly name of the maintenance window."},
			{Name: "status", Type: proto.ColumnType_INT, Description: "Status of the maintenance window. Possible values are 0 (paused), 1 (active)."},
			{Name: "type", Type: proto.ColumnType_INT, Description: "Type of the maintenance window. Possible values are 1 (Once), 2 (Daily), 3 (Weekly), 4 (Monthly)."},
			{Name: "user", Type: proto.ColumnType_INT, Description: "User of the maintenance window."},
			{Name: "start_time", Type: proto.ColumnType_STRING, Description: "Start time of the maintenance window."},
			{Name: "duration", Type: proto.ColumnType_INT, Description: "Duration of the maintenance windows in minutes."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Separated with '-' and used only for weekly and monthly maintenance windows."},
		},
	}
}

func listMaintenanceWindows(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listMaintenanceWindows", "connection_error", err)
		return nil, err
	}
	var params = uptimerobotapi.GetMWindowParams{
		//set default limit
		Limit: types.Int(50),
	}

	if q, ok := d.KeyColumnQuals["type"]; ok {
		windowType := q.GetInt64Value()
		params.Type = strconv.FormatInt(windowType, 10)
	}

	if q, ok := d.KeyColumnQuals["duration"]; ok {
		duration := q.GetInt64Value()
		params.Duration = strconv.FormatInt(duration, 10)
	}

	count := 0
	for {
		mw, err := conn.MWindow.GetMWindows(params)
		if err != nil {
			plugin.Logger(ctx).Error("listMaintenanceWindows", "api_error", err)
			return nil, err
		}

		for _, element := range mw.MWindows {
			d.StreamListItem(ctx, element)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		count = count + len(mw.MWindows)
		if count >= mw.Pagination.Total {
			break
		}
		params.Offset = count
	}

	return nil, nil
}
