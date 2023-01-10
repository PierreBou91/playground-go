package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	key      string
	value    string
	expected string
}

var testCases = []testCase{
	{"a", "PLf2e", "PLf2e"},
	{"qwerqwer", "]]32222", "]]32222"},
	{"1234es", "plkmjn", "plkmjn"},
}

func TestPut(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Put %q should return %q", tc.key, tc.expected), func(t *testing.T) {
			Put(tc.key, tc.value)
			if store[tc.key] != tc.expected {
				t.Errorf("Expected %v got %v", tc.expected, store[tc.key])
			}
		})
	}
}

func TestGet(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Get %q should return %q", tc.key, tc.expected), func(t *testing.T) {
			tested, err := Get(tc.key)
			if err != nil {
				t.Errorf("Error in function Get: %v", err)
			}
			if tested != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, tested)
			}
		})
	}
}
