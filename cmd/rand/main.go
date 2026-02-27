package main

import (
	"flag"
	"fmt"
	"github.com/simulib/random"
	"github.com/simulib/random/rng/xoshiro256plusplus"
)

func main() {
	// Command line parsing
	dist := flag.String("dist", "Uniform(0.0,1.0)", "distribution")
	N := flag.Int("N", 10, "number of samples")
	seed := flag.Uint64("seed", 1234, "random seed")
	flag.Parse()

	rng := xoshiro256plusplus.New(*seed)
	R := random.New(rng)
	s, es := R.NewContinuousVariate(*dist)
	if es == nil {
		for i := 0; i < *N; i++ {
			fmt.Println(s())
		}
		return
	}
	d, ed := R.NewDiscreteVariate(*dist)
	if ed == nil {
		for i := 0; i < *N; i++ {
			fmt.Println(d())
		}
		return
	}
	fmt.Println("Unable to parse distribution:", *dist)

}
