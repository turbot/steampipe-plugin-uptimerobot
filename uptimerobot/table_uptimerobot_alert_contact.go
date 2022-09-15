package uptimerobot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION
func tableUptimeAlertContact(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_alert_contact",
		Description: "UptimeRobot Alert Contact",
		List: &plugin.ListConfig{
			Hydrate: listAlertContact,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "unique id of the alert contact.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "friendly_name",
				Type:        proto.ColumnType_STRING,
				Description: "Friendly name of alert contact.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_INT,
				Description: "Type of alert contact.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "Status of alert contact.",
			},
			{
				Name:        "value",
				Type:        proto.ColumnType_STRING,
				Description: "The description of alert policy.",
			},
		},
	}
}

//// LIST FUNCTION
func listAlertContact(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	contacts, err := conn.AllAlertContacts()
	if err != nil {
		return nil, err
	}

	for _, item := range contacts {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}
