package main

func main() {

}

func majorityElement(nums []int) int {
	result := 0
	intMax := 0
	m := make(map[int]int)
	for _, v := range nums {
		m[v] += 1
	}
	for k, v := range m {
		if v > intMax {
			intMax = v
			result = k
		}
	}
	return result
}
