package readable

import (
	"fmt"
	"strings"
)

// KeyValue is the default formatter, and formats like so:
//
//  Log("package", "server", "listern", ":3000")
//  //=> 2015/08/21 20:01:48 package=server listener=:3000
//
func KeyValue(parts ...interface{}) (line string) {

	segments := []string{}

	for i := 0; i < len(parts); i++ {
		if i%2 == 0 {
			if i == len(parts)-1 {
				segments = append(segments, fmt.Sprintf("%+v", parts[i]))
			} else {
				segments = append(segments, fmt.Sprintf("%+v=%+v", parts[i], parts[i+1]))
			}
		}

	}

	return strings.Join(segments, " ")
}

// KeyValueQuoted formats like so:
//
//  readable.SetFormatter(readable.KeyValueQuoted)
//  readable.Log("package", "server", "listern", ":3000", "extra stuff")
//  //=> 2015/08/21 20:01:48 package="server" listener=":3000" "extra stuff"
//
func KeyValueQuoted(parts ...interface{}) (line string) {

	segments := []string{}

	for i := 0; i < len(parts); i++ {
		if i%2 == 0 {
			if i == len(parts)-1 {
				val := fmt.Sprintf("%+v", parts[i])
				segments = append(segments, fmt.Sprintf("%q", val))
			} else {
				val := fmt.Sprintf("%+v", parts[i+1])
				segments = append(segments, fmt.Sprintf("%+v=%q", parts[i], val))
			}
		}

	}

	return strings.Join(segments, " ")
}

// Join is a simple formatter which formats like so:
//
//  readable.SetFormatter(readable.Join)
//  readable.Log("package", "server", "listern", ":3000")
//  //=> "2015/08/21 20:01:48 package server listener :3000"
func Join(parts ...interface{}) string {
	segments := []string{}

	for i := 0; i < len(parts); i++ {
		if i%2 == 0 {
			if i == len(parts)-1 {
				segments = append(segments, fmt.Sprintf("%+v", parts[i]))
			} else {
				segments = append(segments, fmt.Sprintf("%+v: %+v", parts[i], parts[i+1]))
			}
		}

	}

	return strings.Join(segments, " ")
}
