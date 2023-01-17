package main

import "testing"

func TestCheckData(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{key: ""},
			want: false,
		},
		{
			args: args{key: "0fd5C343d6Db3d381628FCD25E19E5f2dEbc6Fbb"},
			want: true,
		},
		{
			args: args{key: "00000000219ab540356cbb839cbe05303d7705fb"},
			want: false,
		},
		{
			args: args{key: "69b4af80Bd555475c870d2C1E84A59B50c9ebFB6"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckData(tt.args.key); got != tt.want {
				t.Errorf("CheckData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteTo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateModelFIle()
		})
	}
}

func TestReadFrom(t *testing.T) {
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
