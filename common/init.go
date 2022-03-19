package common

import (
	"encoding/gob"
)

func init() {
	gob.Register(Hash{})

}
