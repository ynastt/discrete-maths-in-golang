//Реализуем алгоритм построения компонент рёберной двусвязности простого графа
//(два обхода графа в глубину)
package main

import "fmt"
import "github.com/skorobogatov/input"

type vertex struct {
        mark string
        parent *vertex
	comp int
	edge []*vertex
	
}

var queue [](*vertex)

func Enqueue(queue [](*vertex), v *vertex) []*vertex {
        return append(queue, v) 
}

func Dequeue(queue [](*vertex), i int) *vertex {
        return queue[i] 
}

func VisitVertex1(v *vertex) {
	v.mark = "gray"
	queue = Enqueue(queue, v)
	for i, _ := range v.edge {
		if v.edge[i].mark == "white" {
			v.edge[i].parent = v
			VisitVertex1(v.edge[i])
		}
	}
	v.mark = "black"
}

func VisitVertex2(v *vertex, component int) {
	v.comp = component
	for i, _ := range v.edge {
		if v.edge[i].comp == -1 && v.edge[i].parent != v {
                        VisitVertex2(v.edge[i], component) 
                }
	}
}

func main() {
        var component1, component2 int
        var n, m int // n- вершин m-ребер графа
	//fmt.Scan(&n, &m)
	input.Scanf("%d %d", &n, &m)
	list := make([]vertex, n)
	for i := 0; i < m; i++ {
                var u, v int //инцидентные ребру вершины
                //fmt.Scan(&u, &v)
	        input.Scanf("%d %d", &u, &v)
                list[u].edge = append(list[u].edge, &list[v])
                list[v].edge = append(list[v].edge, &list[u])
        }
        //DFS1
	for i, _ := range list { 
                list[i].mark = "white" 
        }
	for i, _ := range list {
		if list[i].mark == "white" { 
			VisitVertex1(&list[i])
 			component1++ 
		}
	}
        //DFS2
	for i, _ := range list {
                list[i].comp = -1 
        }
	for i, _ := range queue {
		v := Dequeue(queue, i)
		if v.comp == -1 {
			VisitVertex2(v, component2)
			component2++
		}
	}
        //num
        //fmt.Println()
        bridges := component2-component1
	fmt.Println(bridges)
}