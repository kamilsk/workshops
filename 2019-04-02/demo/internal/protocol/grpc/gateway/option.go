package gateway

import "context"

type Option func(*server) error

func WithMaintenanceHandler(ctx context.Context) Option {
	return func(server *server) error {
		return RegisterMaintenanceHandler(ctx, server.mux, server.conn)
	}
}

func WithSequenceHandler(ctx context.Context) Option {
	return func(server *server) error {
		return RegisterSequenceHandler(ctx, server.mux, server.conn)
	}
}
