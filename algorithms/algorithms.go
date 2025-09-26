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
	fmt.Println("Запуск алгоритма на нахождении суммы двух чисел в срезе")
	slice1 := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("Индексы чисел из среза %v, которые в сумме будут давать %v : %v\n ", slice1, target, Two_sum(slice1, target))
	fmt.Println("Запуск алгоритма на проверку уникального кол-ва повторений значений")
	slice2 := []int{1, 2, 2, 1, 1, 3}
	fmt.Printf("Проверка уникальности повторений для среза %v: %v\n ", slice2, UniqueOccurrences(slice2))
	fmt.Println("Запуск алгоритма на проверку главного числа в слайсе")
	slice3 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("Проверка главного числа для среза %v: %v\n ", slice3, MajorityElement(slice2))
	fmt.Println("Алгоритм проверки скобок:")
	fmt.Println("({[]})", IsBalanced("({[]})"))
	fmt.Println("([)]", IsBalanced("([)]"))
	fmt.Println("{(})", IsBalanced("{(})"))
	fmt.Println("Пустая строка", IsBalanced(""))
	fmt.Println("({[text]})", IsBalanced("({[text]})"))
	fmt.Println("Алгоритм реверса строки Hello World:")
	fmt.Println(ReverseStrings("Hello World!"))
	text := "абвгдейка"
	textTarget := "где"
	fmt.Printf("Алгортим поиска подстроки '%s', в тексте '%s' : \n", textTarget, text)
	fmt.Println(NaiveSearch(text, textTarget))
}
