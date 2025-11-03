package main

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "positive",
			a:        10,
			b:        50,
			expected: 500,
		},
		{
			name:     "negative",
			a:        -10,
			b:        -3,
			expected: 30,
		},
		{
			name:     "mixed",
			a:        -4,
			b:        5,
			expected: -20,
		},
		{
			name:     "zero",
			a:        0,
			b:        5,
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Multiply(tt.a, tt.b)
			if res != tt.expected {
				t.Errorf("Multiply(%d, %d)=%d, expected:%d", tt.a, tt.b, res, tt.expected)
			}
		})
	}
}
