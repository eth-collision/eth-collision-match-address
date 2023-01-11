package main

import "testing"

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
			want: false,
		},
		{
			args: args{address: "0xc0e99e3981e09da3452b5bcb68ee33b6576bd5d7"},
			want: true,
		},
		{
			args: args{address: "0xc0e99e3981e09da3452b5bcb68ee33b6576bd5D7"},
			want: true,
		},
		{
			args: args{address: "0xc0e99e3981e09da3452b5bcb68ee33b6576bd5D8"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkAddress(tt.args.address); got != tt.want {
				t.Errorf("checkAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
