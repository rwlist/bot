package mlog

import (
	"log/slog"
	"os"
)

func Fatal(msg string, attrs ...any) {
	slog.Error(msg, attrs...)
	os.Exit(1)
}
