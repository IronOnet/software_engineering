package main 

func twoSum(nums []int, target int) []int{
	m := make(map[int]int) 
	for index, num := nums{
		if val, ok := m[target-num]; ok{
			return []int{val, index}
		}
		m[num] = index
	}
	return nil 
}