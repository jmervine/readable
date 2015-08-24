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

    logger := readable.New().WithPrefix("logger")
    debugger := readable.New().WithDebug().WithPrefix("debugger")

    logger.Log("type", "logger", "fn", "Log")
    logger.Debug("type", "logger", "fn", "Debug")
    logger.WithDebug().Debug("type", "logger", "fn", "WithDebug().Debug")
    // 2015/08/23 19:17:38 logger type=logger fn=Log
    // 2015/08/23 19:17:38 logger type=logger fn=WithDebug().Debug

    debugger.Log("type", "logger", "fn", "Log")
    debugger.Debug("type", "logger", "fn", "Debug")
    // 2015/08/23 19:17:38 debugger type=logger fn=Log
    // 2015/08/23 19:17:38 debugger type=logger fn=Debug
}
```
