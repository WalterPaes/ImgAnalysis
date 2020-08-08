package core

import (
	"fmt"
	"testing"
)

func TestNewHttpConnector(t *testing.T) {
	http := NewHttpConnector()
	expected := "*core.Http"
	got := fmt.Sprintf("%T", http)
	if expected != got {
		t.Errorf("Was expected '%s' but got '%s'", expected, got)
	}
}
