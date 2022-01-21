package main

import "fmt"
import "math"

func main() {
        var x int
        var array1 []int
        var array2 []int
	fmt.Scan(&x)
        sqrtx := math.Round(math.Sqrt(float64(x)))
        for i := 1; i <= int(sqrtx); i++ {
                if x % i == 0 {
                        array1 = append(array1,x/i)
                        if x / i != i {
                                array2 = append(array2,i)
                        }
                }
        }
        for i := len(array2) -1; i>= 0; i-- {
                array1 = append(array1, array2[i])
        }
	fmt.Println("graph {")
	//выводим вершины
        for _, v := range array1 {
		fmt.Println(v)
	}
	//выводим ребра
        for u := 0; u < len(array1); u++ {
                for v := u+1; v < len(array1); v++ {
                        if array1[u] % array1[v] == 0 {
                                f := true;
                                for w := v - 1; w > u && f; w-- {
                                        if array1[u] % array1[w] == 0 &&
                                           array1[w] % array1[v] == 0 {
                                                f = false;
                                        }
                                }
                                if f && array1[u] != 1 {
                                        fmt.Println(array1[u],"--", array1[v])
                                }
                        }
                }
        }
	fmt.Println("}")
}