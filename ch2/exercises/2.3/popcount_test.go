package popcount

import "testing"

func TestPopCount(t *testing.T) {
	testCases := []struct {
		Input  uint64
		Output int
	}{
		{5, 2},
		{1, 1},
	}

	for _, testCase := range testCases {
		if x := PopCount(testCase.Input); x != testCase.Output {
			t.Fatalf("%v != %v", x, testCase.Output)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x123456789ABCDEF)
	}
}
