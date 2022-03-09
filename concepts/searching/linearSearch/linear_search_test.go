package linearSearch_test

import (
	"github.com/hardiksachan/dsaInGo/concepts/searching/linearSearch"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	nums := []int{23, 45, 1, 2, 8, 19, -3}

	t.Run("return -1 if element is not found", func(t *testing.T) {
		assertIndex(t, linearSearch.LinearSearch(nums, -99), -1)
	})

	t.Run("return index if element is found", func(t *testing.T) {
		assertIndex(t, linearSearch.LinearSearch(nums, 45), 1)
	})

	t.Run("return -1 when array is empty", func(t *testing.T) {
		assertIndex(t, linearSearch.LinearSearch([]int{}, 45), -1)
	})

}

func assertIndex(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
