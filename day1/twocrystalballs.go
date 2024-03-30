package day1

import (
	"math"
)

func TwoCrystalBalls(breaks []bool) int {
	// jmpAmount is breaks.length |> sqrt |> floor
	var jmpAmount = int(math.Floor(math.Sqrt(float64(len(breaks)))))

	// Set i to the first jump
	i := jmpAmount

	// Keep jumping the jump amount until it breaks
	for ; i < len(breaks); i += jmpAmount {
		if breaks[i] { break }
	}

	// Go back to the last known safe location
	i -= jmpAmount

	// Search linearly until it gets to the next jump location,
	// or gets to the end of the list
	for j := 0; j <= jmpAmount && i < len(breaks); j++ {
		if breaks[i] {
			// If it breaks, then return the location of the break
			return i
		}

		i++
	}

	// Return -1 if the location was not found
	return -1
}
