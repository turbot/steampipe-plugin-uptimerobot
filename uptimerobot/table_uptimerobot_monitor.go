package uptimerobot

import (
	"context"
	"strconv"

	"github.com/bigdatasourav/uptimerobotapi"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableUptimerobotMonitor(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_monitor",
		Description: "UptimeRobot Monitor used to monitor websites.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getMonitor,
		},
		List: &plugin.ListConfig{
			Hydrate: listMonitors,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "type",
					Require: plugin.Optional,
				},
				{
					Name:    "status",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the monitor (can be used for monitor-specific requests).",
			},
			{
				Name:        "friendly_name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name of the monitor.",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Description: "the URL/IP of the monitor.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_INT,
				Description: "The status of the monitor. When used with the editMonitor method 0 (to pause) or 1 (to start) can be sent.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_INT,
				Description: "the type of the monitor. See monitor>type parameter in https://uptimerobot.com/api/.",
			},
			{
				Name:        "create_date_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The create time for the monitor.",
				Transform:   transform.FromField("CreateDatetime").Transform(convertTimestamp),
			},
			{
				Name:        "http_username",
				Type:        proto.ColumnType_STRING,
				Description: "It used for password-protected web pages. Available for HTTP and keyword monitoring.",
			},
			{
				Name:        "http_password",
				Type:        proto.ColumnType_STRING,
				Description: "It used for password-protected web pages. Available for HTTP and keyword monitoring.",
			},
			{
				Name:        "interval",
				Type:        proto.ColumnType_INT,
				Description: "The interval for the monitoring check (300 seconds by default).",
			},
			{
				Name:        "is_group_main",
				Type:        proto.ColumnType_BOOL,
				Description: "Specify if the monitor group is main.",
				Transform:   transform.FromField("IsGroupMain").Transform(convertToBool),
			},
			{
				Name:        "keyword_case_type",
				Type:        proto.ColumnType_INT,
				Description: "It used only for Keyword monitoring (monitor>type = 2) if set the keyword value will be checked as case sensitive or case insensitive according the selection. (case sensitive by default).",
			},
			{
				Name:        "keyword_type",
				Type:        proto.ColumnType_STRING,
				Description: "It used only for Keyword monitoring (monitor>type = 2) and shows if the monitor will be flagged as down when the keyword exists or not exists.",
			},
			{
				Name:        "keyword_value",
				Type:        proto.ColumnType_STRING,
				Description: "The value of the keyword.",
			},
			{
				Name:        "monitor_group",
				Type:        proto.ColumnType_INT,
				Description: "Number of the monitor group.",
			},
			{
				Name:        "sub_type",
				Type:        proto.ColumnType_STRING,
				Description: "It used only for Port monitoring (monitor>type = 4) and shows which pre-defined port/service is monitored or if a custom port is monitored.",
			},
			{
				Name:        "timeout",
				Type:        proto.ColumnType_INT,
				Description: "The timeout for the monitoring check.",
			},
			{
				Name:        "alert_contacts",
				Type:        proto.ColumnType_JSON,
				Description: "The alert contacts of the monitor.",
			},
			{
				Name:        "logs",
				Type:        proto.ColumnType_JSON,
				Description: "The logs of the monitor.",
			},
			{
				Name:        "ssl",
				Type:        proto.ColumnType_JSON,
				Description: "The ssl of the monitor.",
				Transform:   transform.FromField("SSL"),
			},
		},
	}
}

func listMonitors(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listMonitor", "connection_error", err)
		return nil, err
	}

	input := uptimerobotapi.GetMonitorsParams{
		//set default limit
		Limit: types.Int(50),
	}

	if q, ok := d.KeyColumnQuals["type"]; ok {
		typeMonitor := q.GetInt64Value()
		input.Types = types.String(strconv.FormatInt(typeMonitor, 10))
	}

	if q, ok := d.KeyColumnQuals["status"]; ok {
		status := q.GetInt64Value()
		input.Statuses = types.String(strconv.FormatInt(status, 10))
	}

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

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
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

	// check if the id is empty
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
