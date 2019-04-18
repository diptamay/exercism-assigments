package isogram

import (
	"testing"

	"github.com/diptamay/exercism/go/completed"
)

func TestIsIsogram(t *testing.T) {
	for _, c := range completed.testCases {
		if completed.IsIsogram(c.input) != c.expected {
			t.Fatalf("FAIL: %s\nWord %q, expected %t, got %t", c.description, c.input, c.expected, !c.expected)
		}

		t.Logf("PASS: Word %q", c.input)
	}
}

func BenchmarkIsIsogram(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, c := range completed.testCases {
			completed.IsIsogram(c.input)
		}

	}
}
