package config

import (
	"context"
	"data-parameter/constant"
	"fmt"
	"log/slog"
	"os"
	"time"
)

type LogHandler struct{}

func (h LogHandler) Enabled(context context.Context, level slog.Level) bool {
	switch level {
	case slog.LevelDebug:
		return true
	case slog.LevelInfo:
		return true
	case slog.LevelWarn:
		return true
	case slog.LevelError:
		return true
	default:
		panic("unreachable")
	}
}

func (h LogHandler) Handle(ctx context.Context, record slog.Record) error {
	message := record.Message

	// Ambil Request ID dari context (jika ada)
	requestID, _ := ctx.Value(constant.RequestID).(string)
	if requestID != "" {
		message = fmt.Sprintf("[Request ID: %s] %s", requestID, message)
	}

	// Tambahkan attribute ke message
	record.Attrs(func(attr slog.Attr) bool {
		message += fmt.Sprintf(" %v", attr)
		return true
	})

	timestamp := record.Time.Format(time.RFC3339)

	switch record.Level {
	case slog.LevelDebug, slog.LevelInfo, slog.LevelWarn:
		fmt.Fprintf(os.Stderr, "[%v] %v %v\n", record.Level, timestamp, message)
	case slog.LevelError:
		fmt.Fprintf(os.Stderr, "%v !!!ERROR!!! %v %v\n", requestID, timestamp, message)
	default:
		panic("unreachable")
	}

	return nil
}

// for advanced users
func (h LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("unimplemented")
}

// for advanced users
func (h LogHandler) WithGroup(name string) slog.Handler {
	panic("unimplemented")
}
