package algorithms

// Поиск числа в отсортированном срезе

func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		middle := low + (high-low)/2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	return -1
}

// Поиск самой первой "плохой" версии

func FirstBadVersion(versions []int) int {
	low, high := 0, len(versions)-1
	var indexBadVersion int

	for low <= high {
		middle := low + (high-low)/2

		if isBadVersion(versions[middle]) {
			indexBadVersion = middle
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return indexBadVersion
}

func isBadVersion(version int) bool {
	if version >= 6 {
		return true
	}
	return false
}

// Квадратный корень числа с округлением в меньшую сторону

func RoundedSqrt(x int) int {
	var low int
	high := x

	for low < high-1 {
		middle := low + (high-low)/2

		if middle*middle == x {
			return middle
		}

		if middle*middle > x {
			high = middle
		} else {
			low = middle
		}
	}

	return low
}
