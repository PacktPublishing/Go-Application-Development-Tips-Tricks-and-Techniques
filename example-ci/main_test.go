package main

import (
	"bytes"
	"testing"
)

func TestPrintHelloPrintsHello(t *testing.T) {
	buf := bytes.Buffer{}
	err := PrintHello(&buf)

	if err != nil {
		t.Fatalf("err is expected nil, got: %s", err)
	}

	out := buf.String()

	if out != "Hello World\n" {
		t.Fatalf(`output is not "Hello World\n"; got: "%s"`, out)
	}
}
