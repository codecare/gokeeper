package commands

import (
	"github.com/codecare/gokeeper/internal/crypt"
	"testing"
)

func Test_parseLengthParameter(t *testing.T) {

	tests := []struct {
		name string
		args []string
		want int
	}{
		{name: "no args", args: []string{"g"}, want: 15},
		{name: "valid arg", args: []string{"g", "  17 "}, want: 17},
		{name: "invalid arg", args: []string{"g", "  bullshit "}, want: 15},
		{name: "out of range too small 1", args: []string{"g", "  -123 "}, want: 8},
		{name: "out of range too small 2", args: []string{"g", "  4 "}, want: 8},
		{name: "out of range too big", args: []string{"g", "  160 "}, want: 128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLengthParameter(tt.args); got != tt.want {
				t.Errorf("parseLengthParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_determineAllowedChars(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "no args", args: []string{"g"}, want: crypt.AllowedCharsAll},
		{name: "no args", args: []string{"g", "12"}, want: crypt.AllowedCharsAll},
		{name: "no args", args: []string{"g", "12", ""}, want: crypt.AllowedCharsAll},
		{name: "no args", args: []string{"g", "12", " bullshit"}, want: crypt.AllowedCharsAll},
		{name: "no args", args: []string{"g", "12", "sh"}, want: crypt.AllowedCharsShellFriendly},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineAllowedChars(tt.args); got != tt.want {
				t.Errorf("determineAllowedChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
