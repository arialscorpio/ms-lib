package logger

import (
	"os"

	"github.com/arialscorpio/ms-lib/logger/context"
	"github.com/arialscorpio/ms-lib/logger/log/field"
	"github.com/rs/zerolog"
)

type Logger interface {
	Error(err error, fields ...field.Field)
	Info(msg string, fields ...field.Field)
	WithContext(ctx1 context.Context, ctxN ...context.Context) Logger
}

type ConsoleLogger struct {
	zl zerolog.Logger
}

var _ Logger = new(ConsoleLogger)

func New() *ConsoleLogger {
	return &ConsoleLogger{
		zl: zerolog.New(os.Stderr).With().Timestamp().Logger(),
	}
}

func (l *ConsoleLogger) WithContext(lctx0 context.Context, lctxN ...context.Context) Logger {
	lc := l.zl.With()
	lc = lctx0(lc)

	for _, lctx := range lctxN {
		lc = lctx(lc)
	}

	return &ConsoleLogger{zl: lc.Logger()}
}

func (l *ConsoleLogger) Error(err error, fields ...field.Field) {
	event := l.zl.Error().Err(err)
	for _, f := range fields {
		f(event)
	}
	event.Msg("")
}

func (l *ConsoleLogger) Info(msg string, fields ...field.Field) {
	event := l.zl.Info()
	for _, f := range fields {
		f(event)
	}
	event.Msg(msg)
}
