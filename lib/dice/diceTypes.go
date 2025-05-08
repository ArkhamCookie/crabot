package dice

import "github.com/ArkhamCookie/crabot/dice/roll"

// D2 rolls a D2 (1-2) `count` amount of times.
func D2(count int) (result int) {
	for range count {
		result += roll.Roll(2)
	}
	return result
}

// D4 rolls a D4 (1-4) `count` amount of times.
func D4(count int) (result int) {
	for range count {
		result += roll.Roll(4)
	}
	return result
}

// D6 rolls a D6 (1-6) `count` amount of times.
func D6(count int) (result int) {
	for range count {
		result += roll.Roll(6)
	}
	return result
}

// D8 rolls a D8 (1-8) `count` amount of times.
func D8(count int) (result int) {
	for range count {
		result += roll.Roll(8)
	}
	return result
}

// D10 rolls a D10 (1-10) `count` amount of times.
func D10(count int) (result int) {
	for range count {
		result += roll.Roll(10)
	}
	return result
}

// D12 rolls a D12 (1-12) `count` amount of times.
func D12(count int) (result int) {
	for range count {
		result += roll.Roll(12)
	}
	return result
}

// D20 rolls a D20 (1-20) `count` amount of times.
func D20(count int) (result int) {
	for range count {
		result += roll.Roll(20)
	}
	return result
}

// D00 rolls a D00 (percentile dice) `count` amount of times.
//
// Results are stored in `[]int`.
func D00(count int) []int {
	var result int
	// Create the results array
	results := make([]int, count)

	// Set count to the correct count for arraying
	count--
	// Get *count* number of results
	for count >= 0 {
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

		// Mark count as done
		count--
	}

	return results
}
