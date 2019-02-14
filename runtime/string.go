package lib

import (
	"log"
	"wasmgo/types"
)

type String struct {
	types.VMInterface
}

//TODO
func (s *String) Strcat(dest int64, src int64) int64 {
	log.Fatalf("Strcat not implementation! ")
	return 0
}
func (s *String) Strlen(str int64) int64 {
	log.Fatalf("Strlen not implementation! ")
	return 0
}
