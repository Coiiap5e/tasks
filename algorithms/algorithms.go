package algorithms

func RunAlgorithm() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	MergeSort(arr1, 3, arr2, 3)
	s := "A man, a plan, a canal: Panama"
	_ = IsPalindromeUppercase(s)
	s2 := "abcdcbad"
	_ = IsPalindromeWithoutOneLetter(s2)
	_ = TwoSum(arr2, 7)
	_ = UniqueOccurrences(arr2)
	_ = MajorityElement(arr2)
	_ = IsBalanced("({[]})")
	_ = ReverseStrings("Hello World!")
	_ = NaiveSearch(s2, "cba")
}
