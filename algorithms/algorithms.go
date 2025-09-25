package algorithms

import "fmt"

func RunAlgorithm() {
	fmt.Println("Запуск алгоритма слияния двух слайсов")
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	fmt.Printf("Срез №1 %v и №2 %v\n", arr1, arr2)
	MergeSort(arr1, 3, arr2, 3)
	fmt.Println(arr1)
	fmt.Println("Запуск алгоритма на проверку палиндрома c лишними символами и заглавными буквами")
	s := "A man, a plan, a canal: Panama"
	fmt.Printf("Проверка строки %s, %v\n", s, IsPalindromeUppercase(s))
	fmt.Println("Запуск алгоритма на проверку палиндрома с возможностью удалить 1 символ")
	s2 := "abcdcbad"
	fmt.Printf("Проверка строки %s, %v\n", s2, IsPalindromeWithoutOneLetter(s2))
}
