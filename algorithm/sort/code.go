package sort

// QuickSort1 1
func QuickSort1(array []int) []int {
	quickSort1(array, 0, len(array)-1)
	return array
}

func quickSort1(array []int, left int, right int) {
	if left >= right {
		return
	}

	index := partition(array, left, right)
	quickSort1(array, left, index-1)
	quickSort1(array, index+1, right)
}

// 分割定位：每一轮可确定base最终所在的位置，base左边均小于它，右边均大于它
func partition(array []int, left int, right int) int {
	base := array[left]
	index := left
	for i := left + 1; i <= right; i++ {
		//分割定位
		if array[i] < base {
			left++
			array[left], array[i] = array[i], array[left]
		}
	}
	array[index], array[left] = array[left], array[index]
	return left
}

func QuickSort2(array []int) []int {

}

// 双指针，挖坑填数
func quickSort2(array []int, left int, right int) {
	if left >= right {
		return
	}
	i, j, base := left, right, array[base] // i是空位
	for i < j {
		if i < j && array[j] >= base { // 从右往左找到第一个小于base的数
			j--
		}
		if i < j { // 找到了第一个小于base的数
			array[i] = array[j] // 填到空位后,j这个位置空缺
			i++
		}
		if i < j && array[i] < base { // 从左往右找到第一个大于base的数
			i++
		}
		if i < j {
			array[j] = array[i]
			j--
		}
	}
	array[i] = base // 循环结束后i位置空缺,同时该位置最终数据确定
	quickSort2(array, left, i-1)
	quickSort2(array, i+1, right)

}
