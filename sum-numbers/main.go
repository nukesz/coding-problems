package main

func main() {

}

// IsAddUpTo Given a list of numbers and a number k,
// returns whether any two numbers from the list add up to k.
func IsAddUpTo(numbers []int, k int) bool {
	count := len(numbers)
	for i, number := range numbers {
		if number == k {
			return true
		}
		j := i
		for {
			j++
			if j == count {
				break
			}
			if number+numbers[j] == k {
				return true
			}
		}
	}
	return false
}
