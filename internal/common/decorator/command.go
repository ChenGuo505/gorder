package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

// CommandHandler defines a generic type that receives a query Q and returns a result R.
type CommandHandler[C, R any] interface {
	Handle(context.Context, C) (R, error)
}

func AppluCommandDecorators[C, R any](handler CommandHandler[C, R], logger *logrus.Entry, metricsClient MetricsClient) CommandHandler[C, R] {
	return queryLoggingDecorator[C, R]{
		logger: logger,
		base: queryMetricsDecorator[C, R]{
			client: metricsClient,
			base:   handler,
		},
	}
}
