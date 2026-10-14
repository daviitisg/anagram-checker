package main

import "testing"

func TestCompare(t *testing.T) {
    tests := []struct {
        word1, word2 string
        expected     bool
    }{
        {"listen", "silent", true},
        {"hello", "world", false},
        {"triangle", "integral", true},
    }

    for _, tc := range tests {
        m1 := count(tc.word1)
        m2 := count(tc.word2)
        result := compare(m1, m2)

        if result != tc.expected {
            t.Errorf("Failed for '%s' and '%s': got %v, expected %v", tc.word1, tc.word2, result, tc.expected)
        } else {
            t.Logf("SUCCESS: '%s' and '%s' matched expected (%v)", tc.word1, tc.word2, tc.expected)
        }
    }
}