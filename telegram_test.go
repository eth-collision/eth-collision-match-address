package main

import "testing"

func Test_sendMsgText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				text: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendMsgText(tt.args.text)
		})
	}
}
