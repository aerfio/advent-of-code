package secure_container

import (
	"fmt"
	"testing"
)

func TestPassword_CheckAdjacentRule(t *testing.T) {
	tests := []struct {
		p    Password
		want bool
	}{
		{
			p:    Password{1, 2, 3, 4, 5, 6},
			want: false,
		},
		{
			p:    Password{1, 1, 3, 4, 5, 6},
			want: true,
		},
		{
			p:    Password{1, 1, 1, 1, 1, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Password %v -> %t", tt.p, tt.want), func(t *testing.T) {
			if got := tt.p.CheckAdjacentRule(); got != tt.want {
				t.Errorf("CheckAdjacentRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword_SatisfiesNotDecreasingRule(t *testing.T) {
	tests := []struct {
		p    Password
		want bool
	}{
		{
			p:    Password{1, 2, 3, 4, 5, 6},
			want: true,
		},
		{
			p:    Password{1, 1, 1, 1, 1, 1},
			want: true,
		},
		{
			p:    Password{1, 1, 1, 0, 1, 1},
			want: false,
		},
		{
			p:    Password{2, 2, 3, 4, 5, 0},
			want: false,
		},
		{
			p:    Password{2, 7, 2, 2, 2, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Password %v -> %t", tt.p, tt.want), func(t *testing.T) {
			if got := tt.p.SatisfiesNotDecreasingRule(); got != tt.want {
				t.Errorf("SatisfiesNotDecreasingRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
