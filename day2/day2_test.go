package main

import "testing"

func Test_checkPassword(t *testing.T) {
	type args struct {
		password  string
		character string
		position  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				password:  "aaaaaa",
				character: "a",
				position:  6,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				password:  "aaaaaa",
				character: "a",
				position:  7,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPassword(tt.args.password, tt.args.character, tt.args.position); got != tt.want {
				t.Errorf("checkPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
