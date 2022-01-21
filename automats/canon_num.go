//need:обход диаграммы автомата в глубину от начального стостояния
package main

import "fmt"

func DFS(q0 int, delta [][]int, r int, m int, was []bool, newnum []int) int {
        was[q0] = true
        newnum[q0] = r
	r++
	for i := 0; i < m; i++ {
		if !was[delta[q0][i]]{
			r = DFS(delta[q0][i], delta, r, m, was, newnum)
		}
	}
        return r
}

func printautomata(delta_new [][]int, phi_new [][]string, n int, m int) {
        fmt.Printf("%d\n%d\n%d\n", n, m, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", delta_new[i][j])
		}
		fmt.Printf("\n")
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%s ", phi_new[i][j])
		}
		fmt.Printf("\n")
	}	
}

func main() {
        var n, m, q0 int
        count := 0
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
        newnum := make([]int, n) //каноническая нумерация
        was := make([]bool, n) //для dfs
	count = DFS(q0, delta, count, m, was, newnum)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			delta[i][j] = newnum[delta[i][j]]
		}
	}
        delta_new := make([][]int, n)
	phi_new := make([][]string, n)
	for i := 0; i < n; i++ {
		delta_new[newnum[i]] = delta[i]
		phi_new[newnum[i]] = phi[i]
	}
        printautomata(delta_new, phi_new, count, m)
}