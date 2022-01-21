package  main

import "fmt"

func printautomata(delta [][]int, phi [][]string, q0 int, n int, m int) {
        fmt.Printf("digraph {\n\trankdir = LR\n\tdummy [label = \"\", shape = none]\n")
	for i := 0; i < n; i++ { 
                fmt.Printf("\t%d [shape = circle]\n", i) 
        }
	fmt.Printf("\tdummy -> %d\n", q0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ { 
                        fmt.Printf("\t%d -> %d [label = \"%s(%s)\"]\n",
                                   i, delta[i][j], string(97+j), phi[i][j]) 
                }
	}
	fmt.Printf("}")
}

func main() {
	var n, m, q0 int
        fmt.Scan(&n) //количество состояний автомата
	fmt.Scan(&m) //размер входного алфавита 
	fmt.Scan(&q0) //номер начального состояния  
	delta := make([][]int, n) //матрицa переходов delta
        phi := make([][]string, n) // матрицa выходов phi
	for i := 0; i < n; i++ {
		delta[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&delta[i][j])
		}
	}
	for i := 0; i < n; i++ {
		phi[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&phi[i][j])
		}
	}
        printautomata(delta, phi, q0, n, m)
}