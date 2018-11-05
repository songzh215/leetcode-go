package main

import (
	"fmt"
)

func isDuplicate(content []byte, c byte) bool {
	size := len(content)
	counter := make(map[byte]int)
	for i := 0; i < size; i++ {
		if _, ok := counter[content[i]]; ok {
			return true
		} else {
			counter[content[i]] = 1
		}
	}
	if _, ok := counter[c]; ok {
		return true
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	content := []byte(s)
	if len(content) <= 1 {
		return len(content)
	}
	window_size := 1
	for i := 1; i < len(content); i++ {
		if !isDuplicate(content[i-window_size:i], content[i]) {
			window_size++
		}
	}
	return window_size
}

func main() {
	s := "abcabcbb"
	s1 := "pwwkew"
	s2 := "bbbbbb"
	res := lengthOfLongestSubstring(s)
	res1 := lengthOfLongestSubstring(s1)
	res2 := lengthOfLongestSubstring(s2)
	fmt.Printf("lengthOfLongestSubstring %s is %d \n", s, res)
	fmt.Printf("lengthOfLongestSubstring %s is %d \n", s1, res1)
	fmt.Printf("lengthOfLongestSubstring %s is %d \n", s2, res2)
}
