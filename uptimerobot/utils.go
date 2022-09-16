package uptimerobot

import (
	"context"
	"math"
	"time"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func convertTimestamp(_ context.Context, d *transform.TransformData) (interface{}, error) {
	epochTime := types.Int(d.Value.(int))

	if epochTime != nil {
		timeInSec := math.Floor(float64(*epochTime) / 1000)
		unixTimestamp := time.Unix(int64(timeInSec), 0)
		timestampRFC3339Format := unixTimestamp.Format(time.RFC3339)
		return timestampRFC3339Format, nil
	}
	return nil, nil
}

func convertToBool(_ context.Context, d *transform.TransformData) (interface{}, error) {
	temp := d.Value.(int)

	if temp == 0 {
		return false, nil
	}
	return true, nil
}
