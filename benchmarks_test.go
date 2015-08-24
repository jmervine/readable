package readable

import (
	"bytes"
	"log"
	"testing"
)

func BenchmarkGolang_Logger(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := new(bytes.Buffer)
		l := log.New(b, "", log.LstdFlags)
		l.Print("foo bar")
	}
}

func BenchmarkReadable_KeyValue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := new(bytes.Buffer)
		r := New().WithOutput(b)
		r.Log("foo", "bar")
	}
}

func BenchmarkReadable_Join(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := new(bytes.Buffer)
		r := New().WithOutput(b).WithFormatter(Join)
		r.Log("foo", "bar")
	}
}
