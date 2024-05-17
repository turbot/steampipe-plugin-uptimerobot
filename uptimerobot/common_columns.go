package uptimerobot

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/uptimerobotapi"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "user_id",
			Description: "User ID of the account.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getAccountUserId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getAccountUserMemoized = plugin.HydrateFunc(getAccountUncached).Memoize(memoize.WithCacheKeyFunction(getAccountUserIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getAccountUserId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	account, err := getAccountUserMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}
	return account.(uptimerobotapi.Account).UserId, err
}

// Build a cache key for the call to getAccountUserIdCacheKey.
func getAccountUserIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getAccountUserId"
	return key, nil
}

func getAccountUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

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

	return account.Account, nil
}
