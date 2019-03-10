// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// HTTP request path constructors for the text service.
//
// Command:
// $ goa gen goa.design/examples/encodings/design -o
// $(GOPATH)/src/goa.design/examples/encodings

package client

import (
	"fmt"
)

// ConcatstringsTextPath returns the URL path to the text service concatstrings HTTP endpoint.
func ConcatstringsTextPath(a string, b string) string {
	return fmt.Sprintf("/concatstrings/%v/%v", a, b)
}

// ConcatbytesTextPath returns the URL path to the text service concatbytes HTTP endpoint.
func ConcatbytesTextPath(a string, b string) string {
	return fmt.Sprintf("/concatbytes/%v/%v", a, b)
}

// ConcatstringfieldTextPath returns the URL path to the text service concatstringfield HTTP endpoint.
func ConcatstringfieldTextPath(a string, b string) string {
	return fmt.Sprintf("/concatstringfield/%v/%v", a, b)
}

// ConcatbytesfieldTextPath returns the URL path to the text service concatbytesfield HTTP endpoint.
func ConcatbytesfieldTextPath(a string, b string) string {
	return fmt.Sprintf("/concatbytesfield/%v/%v", a, b)
}
