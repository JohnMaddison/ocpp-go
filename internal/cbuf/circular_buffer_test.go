package cbuf

import "testing"

func TestFindReturnsMatchingItem(t *testing.T) {
	buf := NewCircularBuffer[int](3)
	buf.Add(10)
	buf.Add(20)
	buf.Add(30)

	item, ok := buf.Find(func(v int) bool { return v == 20 })
	if !ok {
		t.Fatal("expected item to be found")
	}
	if item != 20 {
		t.Fatalf("expected 20, got %d", item)
	}
}

func TestRemoveWhereRemovesFirstMatchingItem(t *testing.T) {
	buf := NewCircularBuffer[int](5)
	for _, item := range []int{1, 2, 3, 4, 5} {
		buf.Add(item)
	}

	item, ok := buf.RemoveWhere(func(v int) bool { return v%2 == 0 })
	if !ok {
		t.Fatal("expected item to be removed")
	}
	if item != 2 {
		t.Fatalf("expected removed item 2, got %d", item)
	}

	got := buf.GetAll()
	want := []int{1, 3, 4, 5}
	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, got)
		}
	}
}

func TestRemoveWhereHandlesWrappedBuffer(t *testing.T) {
	buf := NewCircularBuffer[int](3)
	for _, item := range []int{1, 2, 3, 4} {
		buf.Add(item)
	}

	item, ok := buf.RemoveWhere(func(v int) bool { return v == 3 })
	if !ok {
		t.Fatal("expected item to be removed")
	}
	if item != 3 {
		t.Fatalf("expected removed item 3, got %d", item)
	}

	got := buf.GetAll()
	want := []int{2, 4}
	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, got)
		}
	}
}

func TestRemoveWhereReturnsFalseWhenMissing(t *testing.T) {
	buf := NewCircularBuffer[int](3)
	buf.Add(1)
	buf.Add(2)

	item, ok := buf.RemoveWhere(func(v int) bool { return v == 3 })
	if ok {
		t.Fatalf("expected no removal, got item %d", item)
	}

	got := buf.GetAll()
	want := []int{1, 2}
	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, got)
		}
	}
}
