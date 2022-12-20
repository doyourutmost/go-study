package main

import (
	"fmt"
)

func lengthOfNonRepeatingSubStr(s string) (maxLength int) {
	lastOccurred := make(map[rune]int)
	start := 0
	for i, ch := range s {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return
}
func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) // m2 == empty
	var m3 map[string]int      // m3 == nil
	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key dose not exist")
	}

	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"),
		lengthOfNonRepeatingSubStr("bbbbbb"),
		lengthOfNonRepeatingSubStr("pwwkew"),
		lengthOfNonRepeatingSubStr(""),
		lengthOfNonRepeatingSubStr("b"),
		lengthOfNonRepeatingSubStr("abcd"))
}
