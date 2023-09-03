//Prime number locator, a golang tutorial by Andrew Alston
//David Gatheru

package main

import (
	"flag"
	"fmt"
	"time"
)

func Sieve(max int) []int {
	//Initialize this as an array of uint8's to save memory and save a loop
	//needed to flip the array, this will be initialized as everything 0

	intArray := make([]uint8, max+1)

	//Start with 2, and keep going until the number being tested squared is greater than the maximum number
	for p := 2; p*p <= max; p++ {
		//Check if the array referenced by the prime we are testing is unchanged, if the number is prime
		if intArray[p] == 0 {
			//Change all multiples of the number to state that they aren't prime numbers
			for i := p*2;i <= max; i += p {
				intArray[i] = 1
				}
			}
		}
		//Iterate through the array adding anything that is a prime to our results array
		var primes []int
		for p := 2; p <= max; p++{ 
			if intArray[p] == 0 {
			primes = append(primes, p)
		}
	}
	//Return the resulting array
	return primes
}

func main(){
	var StartTime time.Time

	Maximum := flag.Int("max", 4000, "Maximum number in range to search for primes")
	DumpPrimes := flag.Bool("dump-prime", false, "Dump the list of prime numbers located")
	TimeExecution := flag.Bool("timing", false, "Dump the execution and prime the results")
	flag.Parse()

	//Note: flags are always pointers, so we have to de-reference them, hence the asterix
	if *Maximum < 2 {
		fmt.Printf("Maximum must be a number larger than 1\n")
		flag.PrintDefaults()
		return
	}

	//If our timing flag is set - grab the start time before we start looking for the prime numbers
	if *TimeExecution {
		StartTime = time.Now()
	}

	//Find the prime numbers calling our FindPrime function with our minimum and maximum arguments
	results := Sieve(*Maximum)

	//If our timing flag is set - find the time its taken to find all our prime numbers
	if *TimeExecution {
		fmt.Printf("Took us %s to find all primes in a range of %d numbers\n",time.Since(StartTime), *Maximum -2)
	}
	fmt.Printf("Found %d prime numbers between 2 and %d\n", len(results), *Maximum)

	//If our dump flag is set - dump the list of located prime numbers
	if *DumpPrimes {
		fmt.Printf("Located the following prime numbers in the range 2 -> %d\n", *Maximum)
		for i, p := range results {
			fmt.Printf("\t[%d]: %d\n", i, p)
		}
	}
}

