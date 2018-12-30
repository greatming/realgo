package toml

import (
	"github.com/BurntSushi/toml"
)

var (
	Decode       = toml.Decode
	DecodeFile   = toml.DecodeFile
	DecodeReader = toml.DecodeReader
	Unmarshal    = toml.Unmarshal
)

