package uptimerobot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableUptimeRobotAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_account",
		Description: "Retrieve information about your current account.",
		List: &plugin.ListConfig{
			Hydrate: listAccounts,
		},
		Columns: []*plugin.Column{
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The account e-mail."},
			{Name: "user_id", Type: proto.ColumnType_INT, Description: "User ID of the account."},
			{Name: "first_name", Type: proto.ColumnType_STRING, Description: "First name with which the account is registered."},
			{Name: "sms_credits", Type: proto.ColumnType_INT, Description: "The SMS credits limit of the account."},
			{Name: "monitor_limit", Type: proto.ColumnType_INT, Description: "The max number of monitors that can be created for the account."},
			{Name: "monitor_interval", Type: proto.ColumnType_INT, Description: "The min monitoring interval (in seconds) supported by the account."},
			{Name: "up_monitors", Type: proto.ColumnType_INT, Description: "The number of up monitors."},
			{Name: "down_monitors", Type: proto.ColumnType_INT, Description: "The number of down monitors."},
			{Name: "paused_monitors", Type: proto.ColumnType_INT, Description: "The number of paused monitors."},
			{Name: "total_monitors_count", Type: proto.ColumnType_INT, Description: "Total number of monitors running."},
		},
	}
}

func listAccounts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listAccounts", "connection_error", err)
		return nil, err
	}

	account, err := conn.Account.GetAccountDetails()
	if err != nil {
		plugin.Logger(ctx).Error("listAccounts", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, account.Account)
	return nil, nil
}
