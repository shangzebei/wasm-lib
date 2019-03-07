package wasm

import (
	"testing"
)

func TestPlugin(t *testing.T) {
	aa := PlugInstants("NETWORK")
	aa.Call("Hello")
}

func TestLoadSystem(t *testing.T) {

}
