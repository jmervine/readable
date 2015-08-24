// readable is a simple logger based loosely on the 12 Factor Logs
// (http://12factor.net/logs).
//
// Examples:
//
//  import (
//      "github.com/jmervine/readable"
//  )
//
//  func main() {
//      readable.SetPrefix("server")
//      readable.Log("listener", ":3000")
//      //=> 2015/08/21 20:01:48 server listener=:3000
//
//      logger   := readable.New().WithPrefix("logger").WithFlags(0)
//      debugger := logger.WithDebug().WithPrefix("debug")
//
//      logger.Log("type", "log")
//      logger.Debug("type", "debug")
//      //=> logger type=log
//
//      debugger.Log("type", "log")
//      debugger.Debug("type", "debug")
//      //=> debugger type=log
//      //=> debugger type=debug
//  }
package readable

import (
	"fmt"
	"io"
	"log"
	"os"
)

// create a new log.Logger
func logLogger() *log.Logger {
	return log.New(os.Stderr, "", log.LstdFlags)
}

// default Readable
var std = New()

type Readable struct {
	logger    *log.Logger
	prefix    interface{}
	formatter func(...interface{}) string
	debug     bool
}

// New creates a new Readable
func New() *Readable {
	r := new(Readable)
	r.ensure()

	return r
}

// public:

// WithDebug returns a new Readable with debug set to true.
//
// Example:
//
//  log := New()
//  log.WithDebug().Debug("foo", "bar")
func (r *Readable) WithDebug() *Readable {
	n := r.clone()
	n.debug = true
	return n
}

// SetDebug allows for turning debugging on or off.
func (r *Readable) SetDebug(d bool) {
	r.debug = d
}

// SetDebug allows for turning debugging on or off for the default logger.
func SetDebug(d bool) {
	std.SetDebug(d)
}

// WithFormatter returns a new Readable with the passed formatter function.
func (r *Readable) WithFormatter(f func(...interface{}) string) *Readable {
	n := r.clone()
	n.SetFormatter(f)
	return n
}

// SetFormatter sets the formatter function.
func (r *Readable) SetFormatter(f func(...interface{}) string) *Readable {
	r.formatter = f
	return r
}

// SetFormatter sets the formatter function for the default logger.
func SetFormatter(f func(...interface{}) string) {
	std.SetFormatter(f)
}

// WithPrefix returns a new Readable with the passed prefix.
func (r *Readable) WithPrefix(p interface{}) *Readable {
	n := r.clone()
	n.SetPrefix(p)
	return n
}

// SetPrefix sets the prefix string.
func (r *Readable) SetPrefix(p interface{}) *Readable {
	r.prefix = p
	return r
}

// SetPrefix sets the prefix string for the default logger.
func SetPrefix(p interface{}) {
	std.SetPrefix(p)
}

// WithOutput returns a new Readable with the passed output set.
func (r *Readable) WithOutput(w io.Writer) *Readable {
	n := r.clone()
	n.SetOutput(w)
	return n
}

// SetOutput sets the output writer.
func (r *Readable) SetOutput(w io.Writer) *Readable {
	r.logger.SetOutput(w)
	return r
}

// SetOutput sets the output writer for the default logger.
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// WithFlags returns a new Readable with the passed log.Logger flags set,
// see log package for details.
func (r *Readable) WithFlags(f int) *Readable {
	n := r.clone()
	n.SetFlags(f)
	return n
}

// SetFlags sets log.Logger flags, see log package for details.
func (r *Readable) SetFlags(f int) *Readable {
	r.logger.SetFlags(f)
	return r
}

// SetFlags sets log.Logger flags for the default logger, see log package
// for details.
func SetFlags(f int) {
	std.SetFlags(f)
}

// Reset resets the logger, formatter and prefix values to defaults.
func (r *Readable) Reset() {
	r.logger = nil
	r.formatter = nil
	r.ensure()
	r.prefix = ""
}

// Reset resets the default logger's logger, formatter and prefix values
// to defaults.
func Reset() {
	std.Reset()
}

// Log formats and calls log.Print.
func (r *Readable) Log(parts ...interface{}) {
	r.logger.Print(r.prepLine(parts...))
}

// Log formats and calls log.Print using the default logger.
func Log(parts ...interface{}) {
	std.Log(parts...)
}

// Fatal formats and calls log.Fatal.
func (r *Readable) Fatal(parts ...interface{}) {
	r.logger.Fatal(r.prepLine(parts...))
}

// Fatal formats and calls log.Fatal using the default logger.
func Fatal(parts ...interface{}) {
	std.Fatal(parts...)
}

// Debug formats and calls log.Print when debug is true.
func (r *Readable) Debug(parts ...interface{}) {
	if !r.debug {
		return
	}
	r.Log(parts...)
}

// Debug formats and calls log.Print when debug is true, using the default
// logger.
func Debug(parts ...interface{}) {
	if !std.debug {
		return
	}
	std.Log(parts...)
}

// private:

// clone clones current *Readable
func (r *Readable) clone() *Readable {
	n := new(Readable)

	// force copy readable
	*n = *r

	// force copy logger
	l := *r.logger
	n.logger = &l
	return n
}

// ensure ensures that required values are set
func (r *Readable) ensure() {
	if r.logger == nil {
		r.logger = logLogger()
	}

	if r.formatter == nil {
		r.formatter = KeyValue
	}
}

// prepLine sets up the output string using the current formatter
func (r *Readable) prepLine(parts ...interface{}) string {
	r.ensure()

	var output string
	if r.prefix == nil || r.prefix == "" {
		output = fmt.Sprintf("%s", r.formatter(parts...))
	} else {
		output = fmt.Sprintf("%+v %s", r.prefix, r.formatter(parts...))
	}

	return output
}
