package runner_test

import (
	"1brc/runner"
	"math"
	"strconv"
	"testing"
)

var table []struct {
	line     []byte
	idx      int
	expected float64
}

func init() {
	table = []struct {
		line     []byte
		idx      int
		expected float64
	}{
		{
			[]byte("abc,123.456"),
			3,
			123.456,
		},
		{
			[]byte("abc,-123.456"),
			3,
			-123.456,
		},
		{
			[]byte("abc,0.0"),
			3,
			0.0,
		},
		{
			[]byte("abc,123.1234567891"),
			3,
			123.1234567891,
		},
		{
			[]byte("abc,39.01356467270101"),
			3,
			39.01356467270101,
		},
	}
}
func TestBufToFloat(t *testing.T) {
	for _, test := range table {
		actual := runner.BufToFloat(test.line, test.idx)
		if math.Abs(actual-test.expected) > 0.0000001 {
			t.Fatalf("expected %v got %v", test.expected, actual)
		}
	}
}

func BenchmarkParseFloat(b *testing.B) {
	for range b.N {
		strconv.ParseFloat("123.321", 32)
	}
}

func BenchmarkBufToFloat(b *testing.B) {
	buf := []byte("abc,123.321")
	for range b.N {
		runner.BufToFloat(buf, 3)
	}
}
