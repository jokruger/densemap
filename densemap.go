package densemap

import "fmt"

type Integer interface {
	~int8 | ~int16 | ~int32 | ~uint8 | ~uint16 | ~uint32
}

// DenseMap provides a generic, ID-based lookup structure optimized for fast, contiguous access to values of type T by integer IDs.
type DenseMap[ID Integer, T any] struct {
	minID  ID
	maxID  ID
	values []T
	exists []bool
	count  int
}

// New creates a new DenseMap with a range from minID to maxID (inclusive).
func New[ID Integer, T any](minID, maxID ID) *DenseMap[ID, T] {
	if minID > maxID {
		minID, maxID = maxID, minID
	}
	size := int(maxID - minID + 1)
	return &DenseMap[ID, T]{
		minID:  minID,
		maxID:  maxID,
		values: make([]T, size),
		exists: make([]bool, size),
	}
}

// GetMinID returns the minimum ID in the DenseMap.
func (dm *DenseMap[ID, T]) GetMinID() ID {
	return dm.minID
}

// GetMaxID returns the maximum ID in the DenseMap.
func (dm *DenseMap[ID, T]) GetMaxID() ID {
	return dm.maxID
}

// Set stores a value for the given ID and marks it as valid.
func (dm *DenseMap[ID, T]) Set(id ID, value T) error {
	if id < dm.minID || id > dm.maxID {
		return fmt.Errorf("ID %v out of range [%v, %v]", id, dm.minID, dm.maxID)
	}
	offset := int(id - dm.minID)
	if !dm.exists[offset] {
		dm.count++
	}
	dm.values[offset] = value
	dm.exists[offset] = true
	return nil
}

// Delete removes the value associated with the ID and marks it as invalid.
func (dm *DenseMap[ID, T]) Delete(id ID) error {
	if id < dm.minID || id > dm.maxID {
		return fmt.Errorf("ID %v out of range [%v, %v]", id, dm.minID, dm.maxID)
	}
	offset := int(id - dm.minID)
	if dm.exists[offset] {
		dm.exists[offset] = false
		var zero T
		dm.values[offset] = zero
		dm.count--
	}
	return nil
}

// Get retrieves the value associated with the ID. Returns (value, true) if set, otherwise (zero, false).
func (dm *DenseMap[ID, T]) Get(id ID) (T, bool) {
	if id < dm.minID || id > dm.maxID {
		var zero T
		return zero, false
	}
	offset := int(id - dm.minID)
	return dm.values[offset], dm.exists[offset]
}

// GetPtr retrieves pointer to the value associated with the ID. Returns *value if set, otherwise nil.
func (dm *DenseMap[ID, T]) GetPtr(id ID) *T {
	if id < dm.minID || id > dm.maxID {
		return nil
	}
	offset := int(id - dm.minID)
	if !dm.exists[offset] {
		return nil
	}
	return &dm.values[offset]
}

// Cap returns the total number of possible values (capacity).
func (dm *DenseMap[ID, T]) Cap() int {
	return len(dm.values)
}

// Len returns the number of actually set elements.
func (dm *DenseMap[ID, T]) Len() int {
	return dm.count
}

// IsEmpty returns true if no elements are set.
func (dm *DenseMap[ID, T]) IsEmpty() bool {
	return dm.count == 0
}

// Contains returns true if value is set for a given ID.
func (dm *DenseMap[ID, T]) Contains(id ID) bool {
	if id < dm.minID || id > dm.maxID {
		return false
	}
	offset := int(id - dm.minID)
	return dm.exists[offset]
}

// ForEach iterates over all set values and applies the provided function to each value.
func (dm *DenseMap[ID, T]) ForEach(fn func(id ID, value T)) {
	for i, exists := range dm.exists {
		if exists {
			id := dm.minID + ID(i)
			fn(id, dm.values[i])
		}
	}
}

// Clear resets the DenseMap, removing all set values.
func (dm *DenseMap[ID, T]) Clear() {
	for i := range dm.exists {
		dm.exists[i] = false
		var zero T
		dm.values[i] = zero
	}
	dm.count = 0
}

// Range iterates over a specified range of IDs, applying the provided function to each set value within that range.
func (dm *DenseMap[ID, T]) Range(min, max ID, fn func(id ID, value T)) {
	if min > max {
		min, max = max, min
	}
	if min < dm.minID {
		min = dm.minID
	}
	if max > dm.maxID {
		max = dm.maxID
	}
	for id := min; id <= max; id++ {
		offset := int(id - dm.minID)
		if dm.exists[offset] {
			fn(id, dm.values[offset])
		}
	}
}

// First returns the first set ID and its associated value, or (zeroID, nil) if none are set.
func (dm *DenseMap[ID, T]) First() (ID, *T) {
	for i, exists := range dm.exists {
		if exists {
			id := dm.minID + ID(i)
			return id, &dm.values[i]
		}
	}
	var zeroID ID
	return zeroID, nil
}

// Last returns the last set ID and its associated value, or (zeroID, nil) if none are set.
func (dm *DenseMap[ID, T]) Last() (ID, *T) {
	for i := len(dm.exists) - 1; i >= 0; i-- {
		if dm.exists[i] {
			id := dm.minID + ID(i)
			return id, &dm.values[i]
		}
	}
	var zeroID ID
	return zeroID, nil
}
