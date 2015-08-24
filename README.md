# readable

`readable` is a simple Golang logger based loosely on the 12 Factor Logs (http://12factor.net/logs)

[![GoDoc](https://godoc.org/gopkg.in/jmervine/readable.v1?status.png)](https://godoc.org/gopkg.in/jmervine/readable.v1) [![Build Status](https://travis-ci.org/jmervine/readable.svg?branch=master)](https://travis-ci.org/jmervine/readable)

### NOTE: `readable` currently only supports Go 1.5+

## usage

```go
package main

import "gopkg.in/jmervine/readable.v1"

func main() {
    readable.Log("type", "default", "fn", "Log")
    // 2015/08/23 19:17:38 type=default fn=Log

    readable.SetFlags(0)
    readable.Log("type", "without data stamp", "fn", "Log")
    // type=without data stamp fn=Log

    readable.SetFormatter(readable.Join)
    readable.Log("type", "with Join formatter", "fn", "Log")
    // type: with Join formatter fn: Log

    logger := readable.New().WithPrefix("[INFO]:")
    debugger := readable.New().WithDebug().WithPrefix("[DEBUG]:")

    logger.Log("type", "logger", "fn", "Log")
    logger.Debug("type", "logger", "fn", "Debug")
    logger.WithDebug().Debug("type", "logger", "fn", "WithDebug().Debug")
    // 2015/08/23 19:17:38 [INFO]: type=logger fn=Log
    // 2015/08/23 19:17:38 [INFO]: type=logger fn=WithDebug().Debug

    debugger.Log("type", "debugger", "fn", "Log")
    debugger.Debug("type", "debuger", "fn", "Debug")
    // 2015/08/23 19:17:38 [DEBUG]: type=debugger fn=Log
    // 2015/08/23 19:17:38 [DEBUG]: type=debugger fn=Debug
}
```

## performance

`readable` performs at about 1/2 the speed of the default logger, e.g.:

```
go test . -bench=.
PASS
BenchmarkGolang_Logger-4         2000000               846 ns/op
BenchmarkReadable_KeyValue-4     1000000              1776 ns/op
BenchmarkReadable_Join-4         1000000              1887 ns/op
ok      github.com/jmervine/readable    7.229s
```
