package main

import "fmt"
//Для представления «классов»
type status struct { //каждое состояние
	i int
	parent *status
	depth int
	pi *status // массив, каждому состоянию <-> корень класса
	was bool
	delt []*status
}

func find(x *status) *status {
	var root *status
	if x.parent == x {
		root = x
	} else {
		x.parent = find(x.parent)
		root = x.parent
	}
	return root

}

func union(u *status, v *status) {
	rootU := find(u)
	rootV := find(v)
	if rootU.depth < rootV.depth {
		rootU.parent = rootV
	} else {
		rootV.parent = rootU
		if rootU.depth == rootV.depth && rootU != rootV {
			rootU.depth++
		}
	}
	//return
}

func split1(n int, m int, state []status, phi [][]string) int {
	mm := n //общее количество классов
	for i := 0; i < n; i++ {
		state[i].parent = &state[i]
		state[i].depth = 0
	}
	for q1 := 0; q1 < n; q1++ {
		for q2 := q1 + 1; q2 < n; q2++ {
			if find(&state[q1]) != find(&state[q2]) {
				eq := true
				for x := 0; x < m; x++ {
					if phi[q1][x] != phi[q2][x] {
						eq = false
						break
					}
				}
				if eq {
					union(&state[q1], &state[q2])
					mm -= 1
				}
			}
		}
	}
	for q := 0; q < n; q++ {
		state[q].pi = find(&state[q])
	}
	return m
}

func split(n int, m int, state []status, phi [][]string) int {
	mm := n //общее количество классов
	for i := 0; i < n; i++ {
		state[i].parent = &state[i]
		state[i].depth = 0
	}
	for q1 := 0; q1 < n; q1++ {
		for q2 := q1 + 1; q2 < n; q2++ {
			if state[q1].pi == state[q2].pi &&
				find(&state[q1]) != find(&state[q2]) {
				eq := true
				for x := 0; x < m; x++ {
					w1 := state[q1].delt[x]
					w2 := state[q2].delt[x]
					if w1.pi != w2.pi { //!!!
						eq = false
						break
					}
				}
				if eq {
					union(&state[q1], &state[q2])
					mm -= 1
				}
			}
		}
	}
	for q := 0; q < n; q++ {
		state[q].pi = find(&state[q])
	}
	return m
}

func AufenkampHohn(n int, m int, q0 int, delta [][]int, state []status, phi [][]string,  count *int){
	m1 := split1(n, m, state, phi)
	for {
		m2 := split(n, m, state, phi)
		if m1 == m2 {
			break
		}
		m1 = m2
	}
	//
	DFS(q0, delta, state[q0].pi ,count, m, state)
}

func DFS(q0 int, D [][]int, s *status, count *int, m int, stat []status) {
	s.i = *count
	*count++
	s.was = true
	for i := 0; i < m; i++ {
		qq := D[q0][i]
		if !s.delt[i].pi.was {
			DFS(qq, D, s.delt[i].pi, count, m, stat)
		}
	}
}

func printautomata(state []status, phi [][]string, q0 int, n int, m int) {
	fmt.Printf("digraph {\n\trankdir = LR\n\tdummy [label = \"\", shape = none]\n")
	for i := 0; i < n; i++ {
		fmt.Printf("\t%d [shape = circle]\n", i)
	}
	fmt.Printf("\tdummy -> %d\n", q0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
          if state[i].was {
			fmt.Printf("\t%d -> %d [label = \"%s(%s)\"]\n",
				state[i].i, state[i].delt[j].pi.i, string(97+j), phi[i][j])
		   }
        }
	}
	fmt.Printf("}")
}
func main() {
	var n, m, q0 int
	count := 0
	fmt.Scan(&n)  //количество состояний автомата
	fmt.Scan(&m)  //размер входного алфавита
	fmt.Scan(&q0) //номер начального состояния
	delta := make([][]int, n) //матрицa переходов delta
	phi := make([][]string, n) // матрицa выходов phi
	//храним массив состояний
	state := make([]status, n)
	for i := 0; i < n; i++ {
		delta[i] = make([]int, m)
		state[i].delt = make([]*status, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&delta[i][j])
			state[i].delt[j] = &state[delta[i][j]]
		}
	}

	for i := 0; i < n; i++ {
		phi[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&phi[i][j])
		}
	}
	AufenkampHohn( n, m, q0, delta, state, phi, &count)
	//AufenkampHohn(delta, phi)
	printautomata(state,phi, q0, n, m)
}