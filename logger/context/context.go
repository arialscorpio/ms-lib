package context

import "github.com/rs/zerolog"

type Context func(zerolog.Context) zerolog.Context

func Int(k string, v int) Context {
	return func(c zerolog.Context) zerolog.Context {
		return c.Int(k, v)
	}
}

func Str(k string, v string) Context {
	return func(c zerolog.Context) zerolog.Context {
		return c.Str(k, v)
	}
}

func Bool(k string, v bool) Context {
	return func(c zerolog.Context) zerolog.Context {
		return c.Bool(k, v)
	}
}
