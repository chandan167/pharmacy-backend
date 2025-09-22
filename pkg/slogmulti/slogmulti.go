// pkg/slogmulti/multi.go
package slogmulti

import (
	"context"
	"log/slog"
)

// MultiHandler duplicates each Record to all provided handlers.
type MultiHandler struct {
	handlers []slog.Handler
}

func NewMultiHandler(handlers ...slog.Handler) slog.Handler {
	hs := make([]slog.Handler, 0, len(handlers))
	for _, h := range handlers {
		if h != nil {
			hs = append(hs, h)
		}
	}
	return &MultiHandler{handlers: hs}
}

func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (m *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	// r.Clone ensures each handler gets its own copy (avoid shared state).
	for _, h := range m.handlers {
		if h.Enabled(ctx, r.Level) {
			// pass a clone so handlers can't mutate shared state
			if err := h.Handle(ctx, r.Clone()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	nh := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		nh[i] = h.WithAttrs(attrs)
	}
	return &MultiHandler{handlers: nh}
}

func (m *MultiHandler) WithGroup(name string) slog.Handler {
	nh := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		nh[i] = h.WithGroup(name)
	}
	return &MultiHandler{handlers: nh}
}
