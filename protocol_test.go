// go test -timeout 30s -run TestParseSetCommand$ mdtd-cache -v

package main

import (
	"fmt"
	"testing"
)

func TestParseSetCommand(t *testing.T) {
	cmd := &CommandSet{
		Key:   []byte("Foo"),
		Value: []byte("Bar"),
		TTL:   2,
	}

	fmt.Println(cmd.Bytes())

	//r := bytes.NewReader(cmd.Bytes())
}
