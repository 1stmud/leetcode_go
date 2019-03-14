package main

func main() {
	romanToInt("MCMXCIV")
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
