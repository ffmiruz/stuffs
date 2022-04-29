package slices

import (
	"runtime"
)

// I HOLD A LOT OF DATA
type BigData struct {
	data [99999]int64
}

// ExtractElements returns a new slice with only the elements we care about.
// In this case, that's dictated by an index `limiterIndex`.
//
// All other elements of the slice can be discarded.
//
// What concerns regarding memory management / garbage collection do you have
// about this implementation?
//
// Would you implemented it in any other way?
func ExtractElements(sliceOfData []BigData, limiterIndex int) []BigData {
	return sliceOfData[:limiterIndex]
}

func ExtractElements2(sliceOfData []BigData, limiterIndex int) []BigData {
	retData := make([]BigData, limiterIndex)
	return extractElements(sliceOfData, limiterIndex, retData)
}

func extractElements(sliceOfData []BigData, limiterIndex int, retData []BigData) []BigData {
	copy(retData, sliceOfData)
	sliceOfData = nil
	runtime.GC()
	return retData
}
