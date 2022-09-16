package uptimerobot

import (
	"context"

	"github.com/bigdatasourav/uptimerobotapi"
	// "github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	// "github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableUptimeRobotMaintenanceWindow(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_maintenance_window",
		Description: "UptimeRobot Maintenance Window.",
		List: &plugin.ListConfig{
			Hydrate: listMaintenanceWindows,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_INT, Description: "Unique id of the maintenance window."},
			{Name: "user", Type: proto.ColumnType_INT, Description: "."},
			{Name: "type", Type: proto.ColumnType_INT, Description: "Type of maintenance window."},
			{Name: "friendly_name", Type: proto.ColumnType_STRING, Description: "Friendly name of maintenance window."},
			{Name: "start_time", Type: proto.ColumnType_STRING, Description: "Start time of maintenance window."},
			{Name: "duration", Type: proto.ColumnType_INT, Description: "Duration of maintenance window."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Value of maintenance window."},
			{Name: "status", Type: proto.ColumnType_INT, Description: "Status of maintenance window."},
		},
	}
}

func listMaintenanceWindows(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listMaintenanceWindows", "conection_error", err)
		return nil, err
	}

	// //default value
	// var max *int64
	// a := int64(50)
	// max = &a

	// //list function
	// limit := d.QueryContext.Limit
	// if d.QueryContext.Limit != nil {
	// 	if *limit < *max {
	// 		if *limit < 1 {
	// 			max = types.Int64(1)
	// 		} else {
	// 			max = limit
	// 		}
	// 	}
	// }

	var params = uptimerobotapi.GetMWindowParams{}

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
