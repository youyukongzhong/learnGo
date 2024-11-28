package main

import "testing"

// TestSubstr 表格驱动测试
func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"让我们说中文", 6},
		{"一二三三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}
	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("Got %d for input %s; "+
				"Expected %d",
				actual, tt.s, tt.ans)
		}
	}
}

// BenchmarkSubstr 性能测试
func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 8

	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Got %d for input %s; "+
				"Expected %d",
				actual, s, ans)
		}
	}
}
