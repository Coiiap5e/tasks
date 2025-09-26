package algorithms

func MajorityElement(nums []int) int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	var maxCount int
	for key, count := range counter {
		if count > maxCount && count > len(nums)/2 {
			maxCount = key
		}
		//Если предполагается, что главное число одно и всегда есть, то :
		//if count > len(nums)/2 {
		//	maxCount = key
		//  break
		//}
	}
	return maxCount
}
