package main

import "testing"

func Test_validateHeight(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			input: "150cm",
			want:  true,
		},
		{
			input: "193cm",
			want:  true,
		},
		{
			input: "149cm",
			want:  false,
		},
		{
			input: "194cm",
			want:  false,
		},
		{
			input: "59in",
			want:  true,
		},
		{
			input: "76in",
			want:  true,
		},
		{
			input: "58in",
			want:  false,
		},
		{
			input: "77in",
			want:  false,
		},
		{
			input: "77",
			want:  false,
		},
		{
			input: "76incm",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateHeight(tt.input); got != tt.want {
				t.Errorf("validateHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateHairColor(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{input: "#012345", want: true},
		{input: "#678901", want: true},
		{input: "#abcdef", want: true},
		{input: "#12345A", want: false},
		{input: "#12345g", want: false},
		{input: "#12345", want: false},
		{input: "#1234567", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateHairColor(tt.input); got != tt.want {
				t.Errorf("validateHairColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatePID(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{input: "123456789", want: true},
		{input: "012345678", want: true},
		{input: "000000000", want: true},
		{input: "100000001", want: true},
		{input: "23456789", want: false},
		{input: "0123456789", want: false},
		{input: "10000000a", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validatePID(tt.input); got != tt.want {
				t.Errorf("validatePID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateRange(t *testing.T) {
	tests := []struct {
		name  string
		input string
		lower int
		upper int
		want  bool
	}{
		{
			input: "1920",
			lower: 1920,
			upper: 2002,
			want:  true,
		},
		{
			input: "2002",
			lower: 1920,
			upper: 2002,
			want:  true,
		},
		{
			input: "1919",
			lower: 1920,
			upper: 2002,
			want:  false,
		},
		{
			input: "2003",
			lower: 1920,
			upper: 2002,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateRange(tt.input, tt.lower, tt.upper); got != tt.want {
				t.Errorf("validateRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateEyeColor(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{input: "amb", want: true},
		{input: "blu", want: true},
		{input: "brn", want: true},
		{input: "gry", want: true},
		{input: "grn", want: true},
		{input: "hzl", want: true},
		{input: "oth", want: true},
		{input: "amc", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateEyeColor(tt.input); got != tt.want {
				t.Errorf("validateEyeColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
