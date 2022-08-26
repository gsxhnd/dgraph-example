package example

import (
	"testing"
)

var c = NewClient()

func Test_dc_Setup(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.Setup()
		})
	}
}

func Test_dc_DropAll(t *testing.T) {

	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.DropAll()
		})
	}
}
