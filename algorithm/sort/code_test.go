package sort

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	array := []int{5, 1, 3, 6, 2}
	res := QuickSort1(array)
	fmt.Println(res)
}
