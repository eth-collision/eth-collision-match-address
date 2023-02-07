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
			// a exist addr
			args: args{key: "0fd5C343d6Db3d381628FCD25E19E5f2dEbc6Fbb"},
			want: true,
		},
		{
			// a not exist addr
			args: args{key: "00000000219ab540356cbb839cbe05303d7705fb"},
			want: false,
		},
		{
			// may wrong addr
			args: args{key: "69b4af80Bd555475c870d2C1E84A59B50c9ebFB6"},
			want: false,
		},
		{
			// wrong address
			args: args{key: "3B54688fd562b380e169d577B9a6221c3065Ec55"},
			want: true,
		},
		{
			// okx.eth
			args: args{key: "9C538863BED3334A9F455E3EDfAC68886C123AF2"},
			want: true,
		},
		{
			// may wrong
			args: args{key: "0xe42526c0cFd33A893f71bed8CBfC819183dadf2C"},
			want: false,
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
