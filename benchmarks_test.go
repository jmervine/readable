package readable

import (
	"bytes"
	"log"
	"testing"
)

func BenchmarkGolang_Logger(b *testing.B) {
	buf := new(bytes.Buffer)
	l := log.New(buf, "", log.LstdFlags)
	for n := 0; n < b.N; n++ {
		l.Print("foo bar")
	}
}

func BenchmarkReadable_KeyValue(b *testing.B) {
	buf := new(bytes.Buffer)
	r := New().WithOutput(buf)
	for n := 0; n < b.N; n++ {
		r.Log("foo", "bar")
	}
}

func BenchmarkReadable_KeyValue_LogSafe(b *testing.B) {
	buf := new(bytes.Buffer)
	r := New().WithOutput(buf)
	for n := 0; n < b.N; n++ {
		r.LogSafe("foo", "bar")
	}
}

func BenchmarkReadable_Join(b *testing.B) {
	buf := new(bytes.Buffer)
	r := New().WithOutput(buf).WithFormatter(Join)
	for n := 0; n < b.N; n++ {
		r.Log("foo", "bar")
	}
}

func BenchmarkReadable_WithSETTER(b *testing.B) {
	buf := new(bytes.Buffer)
	r := New().WithOutput(buf)
	for n := 0; n < b.N; n++ {
		r.WithPrefix("prefix").Log("foo", "bar")
	}
}

func BenchmarkReadable_TwoWithSETTERs(b *testing.B) {
	buf := new(bytes.Buffer)
	r := New()
	for n := 0; n < b.N; n++ {
		r.WithPrefix("prefix").WithOutput(buf).Log("foo", "bar")
	}
}
