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

    binary.Write(buf, binary.LittleEndian, int(len(c.Key)))
    binary.Write(buf, binary.LittleEndian, c.Value)
    
    binary.Write(buf, binary.LittleEndian, int(len(c.Key)))
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
        fmt.Println("SET")
    }
}