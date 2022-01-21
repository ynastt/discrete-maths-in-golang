package main

import "fmt"

func main() {
        f := func (val1, val2 uint32) uint32 {
                if val2 == 0 { val1*=10 }
		for i:= val2; i!=0; i/=10 {
			val1*=10
		} 
		val1+=val2
		return val1 
	}
	var n uint32
	var b, c, d, e, i, j uint32
	fmt.Scanf("%d", &n)
	var a = make([]uint32, n, n)
	for i = 0; i < n; i ++ {
 		fmt.Scanf("%d", &a[i])
	}
	for i= 0; i< n-1; i++ {
		b = f(a[i], a[i+1])
		c = f(a[i+1], a[i])
		if c > b {
		        a[i], a[i+1] = a[i+1], a[i]
		}
		for j= i; j>0; j-- {
			d = f(a[j-1], a[j])
			e = f(a[j], a[j-1])
			if e > d {
			        a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
	for _, x := range a {
		fmt.Printf("%d", x)
	}
}