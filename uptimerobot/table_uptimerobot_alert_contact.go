package uptimerobot

import (
	"context"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/uptimerobotapi"
)

//// TABLE DEFINITION
func tableUptimeRobotAlertContact(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_alert_contact",
		Description: "Alert Contacts are used to notify for up or down events.",
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
				Description: "The ID of the alert contact.",
			},
			{
				Name:        "friendly_name",
				Type:        proto.ColumnType_STRING,
				Description: "Friendly name of the alert contact.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_INT,
				Description: "The status of the alert contact. Possible values are 0 (not activated), 1 (paused), 2 (active).",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_INT,
				Description: "The type of the alert contact notified. For a list of possible values, please see 'alertcontact>type' parameter in https://uptimerobot.com/api.",
			},
			{
				Name:        "value",
				Type:        proto.ColumnType_STRING,
				Description: "Alert contact's email address, phone, username, url or api key depending on the alert contact type.",
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

	count := 0
	for {
		contacts, err := conn.AlertContact.GetAlertContacts(params)
		if err != nil {
			plugin.Logger(ctx).Error("listAlertContact", "api_error", err)
			return nil, err
		}

		for _, item := range contacts.AlertContacts {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		count = count + len(contacts.AlertContacts)
		if count >= contacts.Total {
			break
		}
		params.Offset = types.Int(count)
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
