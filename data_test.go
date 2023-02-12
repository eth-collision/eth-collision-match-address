package main

import (
	"log"
	"strings"
	"testing"
)

func TestCheckData(t *testing.T) {
	LoadFromModelFile()
	length := GetBloomLength()
	log.Println("bloom length", length)
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exist addr",
			args: args{key: "c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"},
			want: true,
		},
		{
			name: "exist addr",
			args: args{key: strings.ToLower("c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")},
			want: true,
		},
		{
			name: "exist addr",
			args: args{key: strings.ToUpper("c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckDataInBloom(tt.args.key); got != tt.want {
				t.Errorf("CheckData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateModelFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateModelFile()
		})
	}
}

func TestLoadFromModelFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoadFromModelFile()
		})
	}
}

func TestVerifyFromFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerifyFromFile()
		})
	}
}

func TestRealPositiveRate(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{
			name: "test",
			want: 0.0001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RealPositiveRate()
			if got != tt.want {
				t.Errorf("RealPositiveRate() = %v, want %v", got, tt.want)
			}
			t.Log("RealPositiveRate() = ", got)
		})
	}
}
