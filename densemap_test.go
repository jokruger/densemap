package densemap

import (
	"testing"
)

func TestUint16(t *testing.T) {
	a := 1
	b := 10
	c := 20
	d := 30
	idx := New[uint16, int](uint16(a), uint16(d))
	for i := a; i <= b; i++ {
		idx.Set(uint16(i), i)
	}
	for i := c; i <= d; i++ {
		idx.Set(uint16(i), i)
	}

	for i := -10; i < 40; i++ {
		if idx.Cap() != int(d-a+1) {
			t.Errorf("Expected capacity to be %d, got %d", d-a+1, idx.Cap())
		}

		if idx.Len() != (b-a+1)+(d-c+1) {
			t.Errorf("Expected length to be %d, got %d", (b-a+1)+(d-c+1), idx.Len())
		}

		if (i >= a && i <= b) || (i >= c && i <= d) {
			if !idx.Contains(uint16(i)) {
				t.Errorf("Expected to contain %d, but it does not", i)
			}
			if v, ok := idx.Get(uint16(i)); !ok || v != i {
				t.Errorf("Expected to return %d for %d, got %v (ok: %v)", i, i, v, ok)
			}
		} else {
			if idx.Contains(uint16(i)) {
				t.Errorf("Expected to not contain %d, but it does", i)
			}
			if v, ok := idx.Get(uint16(i)); ok || v != 0 {
				t.Errorf("Expected to return zero for %d, got %v (ok: %v)", i, v, ok)
			}
		}
	}
}

func TestInt16(t *testing.T) {
	a := -5
	b := 10
	c := 20
	d := 30
	idx := New[int16, int](int16(a), int16(d))
	for i := a; i <= b; i++ {
		idx.Set(int16(i), i)
	}
	for i := c; i <= d; i++ {
		idx.Set(int16(i), i)
	}

	for i := -10; i < 40; i++ {
		if idx.Cap() != int(d-a+1) {
			t.Errorf("Expected capacity to be %d, got %d", d-a+1, idx.Cap())
		}

		if idx.Len() != (b-a+1)+(d-c+1) {
			t.Errorf("Expected length to be %d, got %d", (b-a+1)+(d-c+1), idx.Len())
		}

		if (i >= a && i <= b) || (i >= c && i <= d) {
			if !idx.Contains(int16(i)) {
				t.Errorf("Expected to contain %d, but it does not", i)
			}
			if v, ok := idx.Get(int16(i)); !ok || v != i {
				t.Errorf("Expected to return %d for %d, got %v (ok: %v)", i, i, v, ok)
			}
		} else {
			if idx.Contains(int16(i)) {
				t.Errorf("Expected to not contain %d, but it does", i)
			}
			if v, ok := idx.Get(int16(i)); ok || v != 0 {
				t.Errorf("Expected to return zero for %d, got %v (ok: %v)", i, v, ok)
			}
		}
	}
}

func TestNew(t *testing.T) {
	dm := New[uint16, string](5, 1)

	if dm.GetMinID() != 1 {
		t.Errorf("Expected min ID to be 1, got %d", dm.GetMinID())
	}

	if dm.GetMaxID() != 5 {
		t.Errorf("Expected max ID to be 5, got %d", dm.GetMaxID())
	}

	if dm.Set(3, "3") != nil {
		t.Errorf("Expected setting ID 3 to succeed")
	}

	if dm.Set(10, "10") == nil {
		t.Errorf("Expected setting ID 10 to fail")
	}

	if dm.GetPtr(10) != nil {
		t.Errorf("Expected getting ID 10 to return nil")
	}

	if dm.GetPtr(1) != nil {
		t.Errorf("Expected getting ID 1 to return nil")
	}

	if dm.GetPtr(3) == nil {
		t.Errorf("Expected getting ID 3 to return non-nil")
	}
}

func TestGet(t *testing.T) {
	dm := New[uint16, string](5, 1)

	if dm.Set(3, "3") != nil {
		t.Errorf("Expected setting ID 3 to succeed")
	}

	v, ok := dm.Get(3)
	if !ok {
		t.Errorf("Expected getting ID 3 to return ok")
	}
	if v != "3" {
		t.Errorf("Expected getting ID 3 to return '3', got '%s'", v)
	}

	_, ok = dm.Get(10)
	if ok {
		t.Errorf("Expected getting ID 10 to return not ok")
	}

	if *dm.GetPtr(3) != "3" {
		t.Errorf("Expected getting ID 3 to return '3', got '%v'", dm.GetPtr(3))
	}
}

func TestForEach(t *testing.T) {
	dm := New[uint16, string](1, 5)

	dm.Set(3, "3")
	dm.Set(5, "5")

	count := 0
	dm.ForEach(func(id uint16, value string) {
		if id != 3 && id != 5 {
			t.Errorf("Expected ID to be 3 or 5, got %d", id)
		}
		if value != "3" && value != "5" {
			t.Errorf("Expected value to be '3' or '5', got '%s'", value)
		}
		count++
	})

	if count != 2 {
		t.Errorf("Expected ForEach to iterate 2 times, got %d", count)
	}
}

func TestAddRemove(t *testing.T) {
	dm := New[uint16, string](1, 5)
	dm.Set(3, "3")
	dm.Set(5, "5")

	if dm.Len() != 2 {
		t.Errorf("Expected length to be 2, got %d", dm.Len())
	}
	if !dm.Contains(3) {
		t.Errorf("Expected to contain ID 3")
	}
	if !dm.Contains(5) {
		t.Errorf("Expected to contain ID 5")
	}
	if dm.Contains(10) {
		t.Errorf("Expected to not contain ID 10")
	}
	if dm.Contains(2) {
		t.Errorf("Expected to not contain ID 2")
	}

	if dm.Delete(10) == nil {
		t.Errorf("Expected deleting ID 10 to fail")
	}

	if dm.Delete(5) != nil {
		t.Errorf("Expected deleting ID 5 to succeed")
	}
	if dm.Contains(5) {
		t.Errorf("Expected to not contain ID 5 after deletion")
	}
	if dm.Len() != 1 {
		t.Errorf("Expected length to be 1 after deletion, got %d", dm.Len())
	}
}

func TestClear(t *testing.T) {
	dm := New[uint16, string](1, 5)
	dm.Set(3, "3")
	dm.Set(5, "5")

	if dm.IsEmpty() {
		t.Errorf("Expected to not be empty after adding elements")
	}

	dm.Clear()
	if dm.Len() != 0 {
		t.Errorf("Expected length to be 0 after clear, got %d", dm.Len())
	}
	if dm.Contains(3) || dm.Contains(5) {
		t.Errorf("Expected to not contain any IDs after clear")
	}

	if !dm.IsEmpty() {
		t.Errorf("Expected to be empty after clear")
	}
}

func TestRange(t *testing.T) {
	dm := New[uint16, string](1, 5)
	dm.Set(3, "3")
	dm.Set(5, "5")

	count := 0
	dm.Range(10, 0, func(id uint16, value string) {
		if id != 3 && id != 5 {
			t.Errorf("Expected ID to be 3 or 5, got %d", id)
		}
		if value != "3" && value != "5" {
			t.Errorf("Expected value to be '3' or '5', got '%s'", value)
		}
		count++
	})

	if count != 2 {
		t.Errorf("Expected Range to iterate 2 times, got %d", count)
	}
}

func TestFirstLast(t *testing.T) {
	dm := New[uint16, string](1, 5)
	dm.Set(3, "3")
	dm.Set(5, "5")

	id, val := dm.First()
	if id != 3 || val == nil || *val != "3" {
		t.Errorf("Expected First to return (3, '3'), got (%d, '%v')", id, val)
	}

	id, val = dm.Last()
	if id != 5 || val == nil || *val != "5" {
		t.Errorf("Expected Last to return (5, '5'), got (%d, '%v')", id, val)
	}

	dm.Clear()
	id, val = dm.First()
	if id != 0 || val != nil {
		t.Errorf("Expected First to return (0, nil) on empty map, got (%d, '%v')", id, val)
	}
	id, val = dm.Last()
	if id != 0 || val != nil {
		t.Errorf("Expected Last to return (0, nil) on empty map, got (%d, '%v')", id, val)
	}
}
