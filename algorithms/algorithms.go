package algorithms

func RunAlgorithm() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	s := "A man, a plan, a canal: Panama"
	s2 := "abcdcbad"

	MergeSort(arr1, 3, arr2, 3)

	_ = IsPalindromeUppercase(s)

	_ = IsPalindromeWithoutOneLetter(s2)

	_ = TwoSum(arr2, 7)

	_ = UniqueOccurrences(arr2)

	_ = MajorityElement(arr2)

	_ = IsBalanced("({[]})")

	_ = ReverseStrings("Hello World!")

	_ = NaiveSearch(s2, "cba")

	words := []string{"hello", "leetcode"}
	_ = IsAlienSorted(words, "hlabcdefgijkmnopqrstuvwxyz")

	_ = IsAnagram("anagram", "nagaram")

	arr3 := []int{1, 0, 0, 2, 0, 3}
	MoveZeroes(arr3)

	_ = RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}
