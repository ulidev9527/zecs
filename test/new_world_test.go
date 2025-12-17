package test

import (
	"testing"
	"zecs/zecs"
)

func TestNewWorld(t *testing.T) {
	world := zecs.NewWorld()
	if world == nil {
		t.Error("expected non-nil world")
	}
}
