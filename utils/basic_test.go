package utils

import (
	"testing"
)

func TestSort(t *testing.T) {

	items := []string{}
	Sort(items)

	tt := []struct {
		desc      string
		values    []string
		resultMin string
		resultMax string
	}{
		{"sort string slice", []string{"zorro", "anna", "franco", "giovanni"}, "anna", "zorro"},
		{"sort string slice", []string{"zorro", "anna", "2franco", "giovanni"}, "2franco", "zorro"},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			Sort(tc.values)
			minValue := tc.values[0]
			maxValue := tc.values[len(tc.values)-1]
			if len(tc.values) > 0 && (minValue != tc.resultMin || maxValue != tc.resultMax) {
				t.Fatalf("expected %v and %v; got %v and %v", tc.resultMin, tc.resultMax, minValue, maxValue)
			}
		})
	}
}

func TestIsSorted(t *testing.T) {
	tt := []struct {
		desc   string
		values []string
		result bool
	}{
		{"check if it's already sorted", []string{"1", "2", "4", "5"}, true},
		{"check if it's not sorted", []string{"1", "2", "5", "4"}, false},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			r := IsSorted(tc.values)
			if r != tc.result {
				t.Fatalf("expected %v; got %v", tc.result, r)
			}
		})
	}
}

func TestSearchItem(t *testing.T) {
	tt := []struct {
		desc       string
		values     []string
		itemVaule  string
		result     int
		errMessage string
	}{
		{"search item that exists in sorted vaules", []string{"ciao", "danny", "foo", "zorro"}, "danny", 1, ""},
		{"search item that not exists in sorted vaules", []string{"ciao", "danny", "foo", "zorro"}, "test", -1, ""},
		{"search item that exists in not sorted vaules", []string{"danny", "ciao", "foo", "zorro"}, "test", -1, "Can't search because items are not sorted"},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := SearchItem(tc.itemVaule, tc.values)
			if err != nil {
				if err.Error() != tc.errMessage {
					t.Fatalf("expected an error %v; got %v", tc.errMessage, err.Error())
				}
			}
			if result != tc.result {
				t.Fatalf("expected result %v, got %v", tc.result, result)
			}
		})
	}
}
