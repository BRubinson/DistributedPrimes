package main

import "math"

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func IsPrimeSlice(values []int) (primes []int) {
	for i := range values {
		if IsPrime(values[i]) {
			primes = append(primes, values[i])
		}
	}
	return
}
