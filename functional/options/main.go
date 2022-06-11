package main

import "fmt"

func main() {
	// create a default logger
	logger := &Logger{}
	fmt.Println(logger)

	// configure logger using variadic options
	logger.SetOptions(Verbose(), Prefix("some_prefix"))
	fmt.Println(logger)
}

type Logger struct {
	verbosity bool
	prefix    string
}

func (l *Logger) SetOptions(options ...Option) {
	for _, option := range options {
		option(l)
	}
}

type Option func(logger *Logger)

func Verbose() Option {
	return func(logger *Logger) {
		logger.verbosity = true
	}
}

func Prefix(prefix string) Option {
	return func(logger *Logger) {
		logger.prefix = prefix
	}
}
