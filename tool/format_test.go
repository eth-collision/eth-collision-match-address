package tool

import (
	"math/big"
	"testing"
)

func TestFormatBigInt(t *testing.T) {
	type args struct {
		n big.Int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				n: *big.NewInt(1000000000000000000),
			},
			want: "1,000,000,000,000,000,000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatBigInt(tt.args.n); got != tt.want {
				t.Errorf("FormatBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
