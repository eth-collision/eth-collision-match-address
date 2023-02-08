package tool

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGetBalanceFromEthScan(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "test",
			args: args{address: ""},
			want: big.NewInt(0),
		},
		{
			name: "test1",
			args: args{address: "0x6D39C4E60dEf1DfC6d09A8FdB5D075e85F0e5F8d"},
			want: big.NewInt(1022255611789767),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBalanceFromEthScan(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBalanceFromEthScan() = %v, want %v", got, tt.want)
			}
		})
	}
}
