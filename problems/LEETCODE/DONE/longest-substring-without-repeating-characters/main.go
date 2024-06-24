package main

import "fmt"

func main() {

	fmt.Println(lengthOfLongestSubstring("aqwertyaa"), 3)

}

func lengthOfLongestSubstring(s string) int {

	var maxLen int
	m := make(map[byte]int)

	start := 0
	for i := 0; i < len(s); i++ {
		currentVal := s[i]
		idx, ok := m[currentVal]

		if ok && idx >= start {

			start = idx + 1
		}
		m[currentVal] = i

		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}

	return maxLen
}

//
//func lengthOfLongestSubstring(s string) int {
//	var maxLen int
//
//	var substrings []string
//	m := make(map[string]int)
//	if len(s) == 0 {
//		return 0
//	}
//	if len(s) == 1 {
//		return 1
//	}
//	var d int
//
//	for i := 0; i < len(s); i++ {
//		currentVal := s[i : i+1]
//		idx, ok := m[currentVal]
//
//		if ok && idx >= d {
//			substrings = append(substrings, s[d:i])
//			d = i
//		}
//
//		if i == len(s)-1 {
//			if !ok {
//				if len(substrings) > 0 {
//					d = len(s[d:i])
//					substrings = append(substrings, s[d:i+1])
//				}
//				substrings = append(substrings, s[d:i+1])
//
//			}
//		}
//		m[currentVal] = i
//	}
//
//	fmt.Println(substrings)
//	for idx, val := range substrings {
//		if maxLen < len(val) {
//			maxLen = len(val)
//		}
//		if len(substrings) == 1 && len(val) == 1 {
//			return len(s[idx+1:])
//		}
//	}
//	if len(substrings) == 0 {
//		return len(s)
//	}
//
//	return maxLen
//}
