package collect

import (
	"context"

	"metrics"
)

type Func func(ctx context.Context, r metrics.Registry, c metrics.CollectPolicy)
