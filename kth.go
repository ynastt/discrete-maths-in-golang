package main

import "fmt"

func main() {
        pow10 := func (val1 uint64) uint64 {
		var z, i uint64
                z=1
		for i=1; i<=val1; i++ {
			z*=10
		} 
		return z 
	}
	
	var k, count, x, d, res uint64
	fmt.Scanf("%d", &k)
	count = 9
	x = 1
	d = 1
	for k >= count {
		k -= count
		x*=10
		d++
		count = 9 * d * x
	}
	res = (pow10(d-1) + k/d) / pow10(d-k%d-1)
	fmt.Printf("%d", res%10)
}
