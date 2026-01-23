package collect

import (
	"context"

	"github.com/alexgaas/metrics"
)

type Func func(ctx context.Context, r metrics.Registry, c metrics.CollectPolicy)
