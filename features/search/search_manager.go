package search

import (
	set "github.com/deckarep/golang-set"
)

type SearchManager struct {
	set   set.Set
	first bool
}

func NewSearchManager() *SearchManager {
	return &SearchManager{
		first: false,
	}
}

func (sm *SearchManager) Add(addition *set.Set) *set.Set {
	// If it will be the first, it becomes the first set
	if !sm.first {
		sm.set = *addition
		sm.first = true
	} else {
		// If will not, it is intersected with the main set
		sm.set = sm.set.Intersect(*addition)
	}
	return &sm.set
}

func (sm *SearchManager) GetSet() *set.Set {
	return &sm.set
}
