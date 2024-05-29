package main

import (
	"basics/confinement"
	safe_maps "basics/safe-map"
)

func main() {
	println("ConcurrentSafeMaps output :-")
	safe_maps.ConcurrentSafeMaps()

	println("\n\nConfinement output :-")
	confinement.Confinement()

	// fan_out_bounded.FanOutBound()
	// drop_pattern.DropPattern()
	// timeout.Timeout()
}
