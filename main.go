package main

import fan_out_bounded "basics/fan-out-bounded"

func main() {
	// println("ConcurrentSafeMaps output :-")
	// safe_maps.ConcurrentSafeMaps()

	// println("\n\nConfinement output :-")
	// confinement.Confinement()

	fan_out_bounded.FanOutBound()
	// drop_pattern.DropPattern()
	// timeout.Timeout()
}
