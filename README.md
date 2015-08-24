# readable

`readable` is a simple Golang logger based loosely on the 12 Factor Logs (http://12factor.net/logs)

[![GoDoc](https://godoc.org/gopkg.in/jmervine/readable.v1?status.png)](https://godoc.org/gopkg.in/jmervine/readable.v1) [![Build Status](https://travis-ci.org/jmervine/readable.svg?branch=master)](https://travis-ci.org/jmervine/readable)

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

`readable` performs a little slower then the default Go logger. `With{*}` setter
convience methods, do impact performance, so keep that in mind where performance
is king.

```
go test . -bench=.
PASS
BenchmarkGolang_Logger-4                 3000000               518 ns/op
BenchmarkReadable_KeyValue-4             1000000              1210 ns/op
BenchmarkReadable_Join-4                 1000000              1213 ns/op
BenchmarkReadable_WithSETTER-4           1000000              1713 ns/op
BenchmarkReadable_TwoWithSETTERs-4       1000000              1883 ns/op
ok      github.com/jmervine/readable    8.267s
```
