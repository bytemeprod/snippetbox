package prettylog

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"

	"github.com/fatih/color"
)

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func NewPrettyHandler(w io.Writer, opts slog.HandlerOptions) *PrettyHandler {
	return &PrettyHandler{
		Handler: slog.NewJSONHandler(w, &opts),
		l:       log.New(w, "", 0),
	}
}

func (h *PrettyHandler) Handle(ctx context.Context, record slog.Record) error {
	level := record.Level.String() + ": "

	switch record.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	attrs := make(map[string]any, record.NumAttrs())
	record.Attrs(func(a slog.Attr) bool {
		attrs[a.Key] = a.Value.Any()
		return true
	})

	fieldsBytes, err := json.MarshalIndent(attrs, "", " ")
	if err != nil {
		return err
	}

	timeStr := record.Time.Format("[15:05:05.000]")
	msg := color.CyanString(record.Message)

	h.l.Println(timeStr, level, msg, color.WhiteString(string(fieldsBytes)))

	return nil
}
