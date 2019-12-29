package utils

import (
	"sort"
)

func BinarySearch(a []string, x string) bool {
	i := sort.Search(len(a), func(i int) bool { return x <= a[i] })
	if i < len(a) && a[i] == x {
		return true
	}

	return false
}
