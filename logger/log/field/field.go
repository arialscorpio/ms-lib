package field

import "github.com/rs/zerolog"

type Field func(*zerolog.Event)

func Int(k string, v int) Field {
	return func(e *zerolog.Event) {
		e.Int(k, v)
	}
}

func Str(k string, v string) Field {
	return func(e *zerolog.Event) {
		e.Str(k, v)
	}
}

func Bool(k string, v bool) Field {
	return func(e *zerolog.Event) {
		e.Bool(k, v)
	}
}
