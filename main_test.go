package main

import (
	"testing"
)

func Test_checkAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{address: ""},
			want: true,
		},
		{
			args: args{address: "0x00000000219ab540356cbb839cbe05303d7705fa"},
			want: true,
		},
		{
			args: args{address: "0x00000000219ab540356cbb839cbe05303d7705fb"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkAddressInBloom(tt.args.address)
			if got != tt.want {
				t.Errorf("checkAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
