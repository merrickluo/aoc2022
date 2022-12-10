package main

import "testing"

func TestPart1(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 7,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 5,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 6,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 10,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 11,
		},
	}

	for _, tt := range testcases {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			p1 := part1(tt.input, 4)

			if p1 != tt.expected {
				t.Errorf("expected p1: %d, got: %d", tt.expected, p1)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 19,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 23,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 23,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 29,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 26,
		},
	}

	for _, tt := range testcases {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			p2 := part1(tt.input, 14)

			if p2 != tt.expected {
				t.Errorf("expected p2: %d, got: %d", tt.expected, p2)
			}
		})
	}
}
