package tool

import (
	"net"
	"testing"
)

func TestGetOutboundIP(t *testing.T) {
	tests := []struct {
		name string
		want net.IP
	}{
		{
			name: "TestGetOutboundIP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetOutboundIP().String()
			t.Log(got)
		})
	}
}
