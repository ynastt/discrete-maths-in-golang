package main

import "fmt"

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) { 
        paritition := func(low, high int,
                           less func(i, j int) bool,
                           swap func(i, j int)) int {
                i, j := low, low
                for j < high {
                        if less(j, high) {
                                swap(i, j)
                                i+=1
                        }
                j++
                }
                swap(i, high)
                return i
        }
        var qsortrec func(low, high int, less func(i, j int) bool)
	qsortrec = func (low, high int, less func(i, j int) bool) {
                if low < high {
                        q := paritition (low, high, less, swap)
                        qsortrec(low, q-1, less)
                        qsortrec(q+1, high, less)
                }
        }
        qsortrec(0, n-1, less)
}

func main() {
	var n, i int
	fmt.Scanf("%d", &n)
	var a = make([]int, n)
	for i = 0; i < n; i ++ {
 		fmt.Scanf("%d", &a[i])
	}
	qsort(n,
		func (i , j int ) bool { return a[i] < a[j] },
		func (i , j int ) { a[i], a[j] = a[j], a[i] })
	for _, x := range a {
		fmt.Printf("%d ", x)
	}
}
