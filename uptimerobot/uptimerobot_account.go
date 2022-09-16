package uptimerobot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tableUptimeRobotAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "uptimerobot_account",
		Description: "UptimeRobot Account.",
		List: &plugin.ListConfig{
			Hydrate: listAccount,
		},
		Columns: []*plugin.Column{
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email with which the account is registered."},
			{Name: "monitor_limit", Type: proto.ColumnType_INT, Description: "Maximum number of monitors allowed for the account."},
			{Name: "monitor_interval", Type: proto.ColumnType_INT, Description: "Monitor interval for the account."},
			{Name: "up_monitors", Type: proto.ColumnType_INT, Description: "Number of monitors up in the account."},
			{Name: "down_monitors", Type: proto.ColumnType_INT, Description: "Number of monitors down in the account."},
			{Name: "paused_monitors", Type: proto.ColumnType_INT, Description: "Number of monitors paused in the account."},
		},
	}
}

func listAccount(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listAccount", "conection_error", err)
		return nil, err
	}
	account, err := conn.Account.GetAccountDetails()
	if err != nil {
		plugin.Logger(ctx).Error("listAccount", "api_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, account.Account)
	return nil, nil
}
