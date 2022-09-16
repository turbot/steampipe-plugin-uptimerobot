package uptimerobot

import (
	"context"

	"github.com/bigdatasourav/uptimerobotapi"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION
func tableUptimeRobotAlertContact(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_alert_contact",
		Description: "UptimeRobot Alert Contact",
		List: &plugin.ListConfig{
			Hydrate: listAlertContact,
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
				// Transform:   transform.FromField("ID"),
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
				Type:        proto.ColumnType_INT,
				Description: "Status of alert contact.",
			},
			{
				Name:        "value",
				Type:        proto.ColumnType_STRING,
				Description: "Value of alert policy.",
			},
		},
	}
}

//// LIST FUNCTION
func listAlertContact(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listAlertContact", "connection_error", err)
		return nil, err
	}

	// Integer type pointer as Limit take *int
	var val *int
	a := 2
	val = &a

	// creating a variable of sruct type
	var params = uptimerobotapi.GetAlertContactsParams{Limit: val}

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
		plugin.Logger(ctx).Error("getAlertPolicy", "connection_error", err)
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetStringValue()

	if id == "" {
		return nil, nil
	}

	params := uptimerobotapi.GetAlertContactsParams{
		AlertContacts: &id,
	}
	result, err := conn.AlertContact.GetAlertContacts(params) // this one is api

	if err != nil {
		plugin.Logger(ctx).Error("getAlertContact", "query_error", err)
		return nil, err
	}

	if len(result.AlertContacts) > 0 {
		return result.AlertContacts[0], nil
	}

	// How to know it will be like this result.AlertContacts[0] ?  ----> inspect the result struct type and found that result {offset, limit, AlertContacts}
	return nil, nil
}
