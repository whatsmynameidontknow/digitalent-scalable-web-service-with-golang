package logging

import (
	"io"
	"log/slog"
)

func New(w io.Writer) *slog.Logger {
	logger := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format("02-Jan-2006 15:04:05 -0700"))
			}
			return a
		},
	}))

	return logger
}
