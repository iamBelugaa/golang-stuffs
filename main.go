package main

import (
	drop_pattern "basics/drop-pattern"
	fan_out_bounded "basics/fan-out-bounded"
	"basics/timeout"
)

func main() {
	// println("ConcurrentSafeMaps output :-")
	// safe_maps.ConcurrentSafeMaps()

	// println("\n\nConfinement output :-")
	// confinement.Confinement()

	fan_out_bounded.FanOutBound()
	drop_pattern.DropPattern()
	timeout.Timeout()
}
