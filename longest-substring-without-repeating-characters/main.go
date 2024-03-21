package main

import (
	"fmt"
)

var s1 = "abcabcbb"
var s2 = "bbbbb"
var s3 = "pwwkew"

func copyInsertCharAt(s string, char rune, i int) string {
	return s[:i] + string(char) + s[i:]
}

func contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring2(s string) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		used := []rune{}
		var ss string
		for j := i; j < len(s); j++ {
			ss = s[i : j+1]
			if !contains(used, rune(s[j])) {
				used = append(used, rune(s[j]))
				if len(ss) > longest {
					longest = len(ss)
				}
			} else {
				break
			}
		}
	}
	return longest
}

func lengthOfLongestSubstring(s string) int {
	longest := 0
	start := 0
	chars := make(map[rune]int)

	for end := 0; end < len(s); end++ {
		char := rune(s[end])
		if _, ok := chars[char]; ok {
			start = max(start, chars[char]+1)
		}
		chars[char] = end
		longest = max(longest, end-start+1)
	}

	return longest
}

func lengthOfLongestSubstring1(s string) int {
	var longest int
	longest = 1
	for i, char := range s {
		fmt.Printf("expanding search starting on '%c'\n", char)
		slen := len(s)
		used := []rune{}
		var ss string
		a, b := i, i
		for {
			if a == 0 && b == slen-1 {
				break
			}
			a_is_dupe := false
			b_is_dupe := false

			if a-1 > -1 {
				a--
			}
			c := rune(s[a])
			fmt.Println(copyInsertCharAt(copyInsertCharAt(s, ']', b), '[', a))
			fmt.Println("a_ss: " + ss)
			if _, err := contains(used, c); err != nil {
				used = append(used, c)
				ss = s[a:b]
				if sslen := len(ss); sslen > longest {
					longest = sslen
				}
			} else {
				a_is_dupe = true
			}
			fmt.Println("a_used: " + string(used))

			if b+1 < slen {
				b++
			}
			c = rune(s[b])
			fmt.Println(copyInsertCharAt(copyInsertCharAt(s, ']', b), '[', a))
			fmt.Println("b_ss: " + ss)
			if _, err := contains(used, c); err != nil {
				used = append(used, c)
				ss = s[a : b+1]
				if sslen := len(ss); sslen > longest {
					longest = sslen
				}
			} else {
				b_is_dupe = true
			}
			fmt.Println("b_used: " + string(used))

			if a_is_dupe && b_is_dupe {
				break
			}
		}
	}
	return longest
}

func main() {
	fmt.Printf(
		"%v len should be 3. actual: %v\n\n",
		s1,
		lengthOfLongestSubstring(s1),
	)

	fmt.Printf(
		"%v len should be 1. actual: %v\n\n",
		s2,
		lengthOfLongestSubstring(s2),
	)

	fmt.Printf(
		"%v len should be 3. actual: %v\n\n",
		s3,
		lengthOfLongestSubstring(s3),
	)

}
