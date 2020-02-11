package utils

import (
	"errors"
	"math"
)

//Sort sorts a slice
func Sort(items []string) {

	if len(items) == 0 {
		return
	}

	index := 1
	outherFlagLoop := true
	insideFlagLoop := true

	for outherFlagLoop {

		outherFlagLoop = false
		insideFlagLoop = true
		index = 1

		for insideFlagLoop {
			prev := items[index-1]
			curr := items[index]

			if prev > curr {
				items[index-1] = curr
				items[index] = prev
				outherFlagLoop = true
			}

			index++

			if index >= len(items) {
				insideFlagLoop = false
			}
		}
	}
}

//IsSorted retrun if a slice of strings is sorted
func IsSorted(items []string) bool {
	for i := 1; i < len(items); i++ {
		if items[i-1] > items[i] {
			return false
		}
	}
	return true
}

//SearchItem searches a string in a slice of strings and return position
//return -1 in case not found or error if not sorted
func SearchItem(itemToSearch string, items []string) (int, error) {
	if IsSorted(items) == false {
		return -1, errors.New("Can't search because items are not sorted")
	}
	return iternalSearch(itemToSearch, items, 0, len(items)), nil
}

func iternalSearch(itemToSearch string, items []string, min int, max int) int {

	candidateIndex := min + int(math.Ceil(float64((max-min)/2)))
	candidateItem := items[candidateIndex]

	if candidateItem == itemToSearch {
		return candidateIndex
	}

	if min > max {
		return -1
	}

	if candidateItem < itemToSearch {
		return iternalSearch(itemToSearch, items, candidateIndex+1, max)
	}

	return iternalSearch(itemToSearch, items, min, candidateIndex-1)
}
