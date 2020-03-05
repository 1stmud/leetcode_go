package main

import "fmt"

// func main() {
// 	// romanToInt("MCMXCIV")
// 	// strs := []string{"flower", "fl", "flight"}
// 	// longestCommonPrefix(strs)
// 	nums := []int{1, 1, 2, 3, 3}

// 	fmt.Println(removeDuplicates(nums))
// }

func longestCommonPrefix(strs []string) string {

	i := 0
	b := false
	var str string
	// var str string
	for {
		var code string
		var waitCompare string
		for k, word := range strs {
			if k == 0 {
				waitCompare = string(word[i])
			} else {
				if len(word) <= i {
					b = true
					break
				}
				code = string(word[i])
				if code != waitCompare {
					b = true
					break
				}
			}
		}
		if b {
			break
		}
		waitCompare = code
		str += waitCompare
		i++
	}
	fmt.Println("str:", str)
	return str
}

func romanToInt(s string) int {

	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	strLen := len(s)
	lastIndex := strLen - 1
	var sum int
	for i := 0; i < lastIndex; i++ {
		current := romanMap[string(s[i])]
		next := romanMap[string(s[i+1])]
		if current < next {
			sum -= current
		} else {
			sum += current
		}
	}
	sum += romanMap[string(s[lastIndex])]
	return sum

}
