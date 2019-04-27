package main

func main() {

}

// IsAddUpTo Given a list of numbers and a number k,
// returns whether any two numbers from the list add up to k.
func IsAddUpTo(numbers []int, k int) bool {
	differences := make(map[int]bool)
	for _, number := range numbers {
		if number == k || differences[number] {
			return true
		}
		differences[k-number] = true
	}
	return false
}
