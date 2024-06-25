package dice

import (
	"fmt"

	"crabot/dice/roll"
)

var (
	result int
)

// D2 rolls a D2 (1-2)
func D2(count int) int {
	for count > 0 {
		result += roll.Roll(2)
		count--
	}
	return result
}

// D4 rolls a D4 (1-4)
func D4(count int) int {
	for count > 0 {
		result += roll.Roll(4)
		count--
	}
	return result
}

// D6 rolls a D6 (1-6)
func D6(count int) int {
	for count > 0 {
		result += roll.Roll(6)
		count--
	}
	return result
}

// D8 rolls a D8 (1-8)
func D8(count int) int {
	for count > 0 {
		result += roll.Roll(8)
		count--
	}
	return result
}

// D10 rolls a D10 (1-10)
func D10(count int) int {
	for count > 0 {
		result += roll.Roll(10)
		count--
	}
	return result
}

// D12 rolls a D12 (1-12)
func D12(count int) int {
	for count > 0 {
		result += roll.Roll(12)
		count--
	}
	return result
}

// D20 rolls a D20 (1-20)
func D20(count int) int {
	for count > 0 {
		result += roll.Roll(20)
		count--
	}
	return result
}

// D00 rolls a D00 (percentile dice)
func D00(count int) []int {
	// Create the results array
	results := make([]int, count)
	fmt.Println(results)

	// Get *count* number of results
	for count > 0 {
		// Roll a D10
		a := roll.Roll(10)

		// Change the D10 into a D00
		if a == 10 {
			a = 0
		} else {
			a = a * 10
		}

		// Roll a D10
		b := roll.Roll(10)

		// If contions for dice were 100 when rolling,
		// add 100 to the results.
		if a == 10 && b == 1 {
			results[count] = 100
			continue
		}

		// Add the results together
		result = a + b
		// Add latest roll to the array
		results[count] = result
	}

	fmt.Println(results)

	return results
}
