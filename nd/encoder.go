package nd

import (
	"encoding/hex"
	"net-debugger/util"
	"strings"
)

type Encoder interface {
	Encode([]byte) []byte
	Decode([]byte) []byte
}

type PlainEncoder struct {
}

func (p PlainEncoder) Encode(bytes []byte) []byte {
	return bytes
}

func (p PlainEncoder) Decode(bytes []byte) []byte {
	return bytes
}

type HexEncoder struct {
}

func (h HexEncoder) Encode(bytes []byte) []byte {
	return []byte(hex.EncodeToString(bytes))
}

func (h HexEncoder) Decode(bytes []byte) []byte {
	s := string(bytes)
	s = strings.ReplaceAll(s, " ", "")
	b, err := hex.DecodeString(s[:len(s)-2])
	util.CheckError(err, "failed to decode hex")
	return b
}
