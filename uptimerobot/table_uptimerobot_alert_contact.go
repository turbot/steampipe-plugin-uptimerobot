package uptimerobot

import (
	"context"

	"github.com/bigdatasourav/uptimerobotapi"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION
func tableUptimeRobotAlertContact(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_alert_contact",
		Description: "UptimeRobot Alert Contact",
		List: &plugin.ListConfig{
			Hydrate: listAlertContacts,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAlertContact,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Unique id of the alert contact.",
			},
			{
				Name:        "friendly_name",
				Type:        proto.ColumnType_STRING,
				Description: "Friendly name of the alert contact.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_INT,
				Description: "Status of the alert contact.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_INT,
				Description: "Type of the alert contact.",
			},
			{
				Name:        "value",
				Type:        proto.ColumnType_STRING,
				Description: "Value of the alert contact.",
			},
		},
	}
}

//// LIST FUNCTION
func listAlertContacts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listAlertContacts", "connection_error", err)
		return nil, err
	}

	// Set default maximum value
	maxVal := types.Int(50)

	// If the requested number of items is less than the paging max limit set the limit to that instead
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < int64(*maxVal) {
			if *limit < 1 {
				maxVal = types.Int(1)
			} else {
				maxVal = types.Int(int(*limit))
			}
		}
	}

	var params = uptimerobotapi.GetAlertContactsParams{Limit: maxVal}

	contacts, err := conn.AlertContact.GetAlertContacts(params)
	if err != nil {
		plugin.Logger(ctx).Error("listAlertContact", "api_error", err)
		return nil, err
	}

	for _, item := range contacts.AlertContacts {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS
func getAlertContact(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getAlertContact", "connection_error", err)
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetStringValue()

	// Check if id is empty.
	if id == "" {
		return nil, nil
	}

	params := uptimerobotapi.GetAlertContactsParams{
		AlertContacts: &id,
	}
	result, err := conn.AlertContact.GetAlertContacts(params)

	if err != nil {
		plugin.Logger(ctx).Error("getAlertContact", "query_error", err)
		return nil, err
	}

	if len(result.AlertContacts) > 0 {
		return result.AlertContacts[0], nil
	}

	return nil, nil
}
