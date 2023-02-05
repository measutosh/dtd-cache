package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Command byte 

const (
    CmdNonce Command = iota
    CmdSet
    CmdGet
    CmdDel
)

type CommandSet struct {
    Key   []byte
    Value []byte
    TTL   int
}

func (c *CommandSet) Bytes () []byte{
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, CmdSet)

    keyLen := int32(len(c.Key))
    binary.Write(buf, binary.LittleEndian, keyLen)
    binary.Write(buf, binary.LittleEndian, c.Value)

    valueLen := int32(len(c.Value))
    binary.Write(buf, binary.LittleEndian, valueLen)
    binary.Write(buf, binary.LittleEndian, c.Value)
    
    binary.Write(buf, binary.LittleEndian, int(c.TTL))

    return buf.Bytes()
}


func ParseCommand(r io.Reader) {
    // Parse commands straight from the connection
    var cmd Command
    binary.Read(r, binary.LittleEndian, &cmd)

    switch cmd {
        case CmdSet:
            set := parseSetCommand(r)
            fmt.Println(set)
    }
}


func parseSetCommand(r io.Reader) *CommandSet {
    cmd := &CommandSet{}

    var keyLen int64
    binary.Read(r, binary.LittleEndian, keyLen)
    cmd.Key = make([]byte, keyLen)
    binary.Read(r, binary.LittleEndian, cmd.Key)

    return cmd
}





