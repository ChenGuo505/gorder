package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

// QueryHandler defines a generic type that receives a query Q and returns a result R.
type QueryHandler[Q, R any] interface {
	Handle(context.Context, Q) (R, error)
}

func AppluQueryDecorators[H, R any](handler QueryHandler[H, R], logger *logrus.Entry, metricsClient MetricsClient) QueryHandler[H, R] {
	return queryLoggingDecorator[H, R]{
		logger: logger,
		base: queryMetricsDecorator[H, R]{
			client: metricsClient,
			base:   handler,
		},
	}
}
