package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
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

func Test_checkAccount(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Println(err)
	}
	type args struct {
		key     *ecdsa.PrivateKey
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				key:     key,
				address: "0x6D39C4E60dEf1DfC6d09A8FdB5D075e85F0e5F8d",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkAccount(tt.args.key, tt.args.address)
		})
	}
}
