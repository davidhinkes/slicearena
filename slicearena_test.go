package slicearena

import (
	"testing"
)

func TestCapacity(t *testing.T) {
	arena := New(0)
	testCap(t, arena, 0)
	testLen(t, arena, 0)
	s := arena.MakeSlice(128).([]int)
	if want, got := 128, len(s); want != got {
		t.Errorf("slice length want %v got %v", want, got)
	}
}

func testCap(t *testing.T, arena *T, want int) {
	t.Helper()
	if got := arena.cap(); want != got {
		t.Errorf("capacity want %v got %v", want, got)
	}
}

func testLen(t *testing.T, arena *T, want int) {
	t.Helper()
	if got := arena.len(); want != got {
		t.Errorf("capacity want %v got %v", want, got)
	}
}

func BenchmarkCalls(b *testing.B) {
	arena := New(0)
	exampleUsage(arena)
	for i := 0; i < b.N; i++ {
		arena.Reset()
		exampleUsage(arena)
	}
}

func exampleUsage(a *T) {
	a.MakeSlice(1024 * 1024)
	a.MakeSlice(22)
	a.MakeSlice(43)
}
