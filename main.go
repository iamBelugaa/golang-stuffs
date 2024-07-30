package main

import (
	token_bucket_algorithm "basics/token-bucket-algorithm"
)

func main() {
	// println("ConcurrentSafeMaps output :-")
	// safe_maps.ConcurrentSafeMaps()

	// println("\n\nConfinement output :-")
	// confinement.Confinement()

	// fan_out_bounded.FanOutBound()
	// drop_pattern.DropPattern()
	// timeout.Timeout()

	token_bucket_algorithm.Run()
}
