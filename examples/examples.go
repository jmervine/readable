package main

import (
	".."
)

func main() {
	readable.Log("type", "default", "fn", "Log")

	readable.SetFlags(0)
	readable.Log("type", "without data stamp", "fn", "Log")

	readable.SetFormatter(readable.Join)
	readable.Log("type", "with Join formatter", "fn", "Log")

	logger := readable.New().WithPrefix("logger")
	debugger := readable.New().WithDebug().WithPrefix("debugger")

	logger.Log("type", "logger", "fn", "Log")
	logger.Debug("type", "logger", "fn", "Debug")
	logger.WithDebug().Debug("type", "logger", "fn", "WithDebug().Debug")

	debugger.Log("type", "logger", "fn", "Log")
	debugger.Debug("type", "logger", "fn", "Debug")
}
