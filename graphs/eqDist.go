package main

import "fmt"
import "sort"//чтобы не писать отдельно функцию сортировки
import "github.com/skorobogatov/input"

type vertex struct {
        mark bool
	num int
	edge []int
}

type Queue struct {
	data []*vertex
	capc int 
	count int 
	head int 
	tail int
}

func InitQueue(queue *Queue, n int) {
        (*queue).data = make([]*vertex, n)
        (*queue).capc = n
        (*queue).count = 0
        (*queue).head = 0
        (*queue).tail = 0  
}

func QueueEmpty(queue *Queue) bool {
        return (*queue).count <= 0
}

func Enqueue(queue *Queue, v *vertex) {
        if (*queue).count == (*queue).capc {
                fmt.Println("Ошибка:переполнение")
        }
        (*queue).data[(*queue).tail] = v
        (*queue).tail += 1
        if (*queue).tail ==  (*queue).capc {
		(*queue).tail = 0
        }
        (*queue).count += 1
}

func Dequeue(queue *Queue) *vertex {
        if QueueEmpty(queue) {
                fmt.Println("Ошибка:опустошение")
        }
        var v *vertex 
        v = (*queue).data[(*queue).head] 
        (*queue).head +=1
        if (*queue).head ==  (*queue).capc {
		(*queue).head = 0
        }
        (*queue).count -= 1
        return v
}

func BFS(a int, queue *Queue, graph []vertex, dist *[]int) {
        for i, _ := range graph {
		graph[i].mark = false
	}
        //для опорной вершины запускаем алгоритм bfs
	graph[a].mark = true
	Enqueue(queue, &graph[a])
	for !QueueEmpty(queue) {
		v := Dequeue(queue)
		for _, x := range (*v).edge {
			if !graph[x].mark {
				(*dist)[x] = (*dist)[v.num] + 1
				graph[x].mark = true
				Enqueue(queue , &graph[x])
			}
		}
	}
}

func main() {
        var n, m, k int 
        var queue Queue
        //fmt.Scan(&n, &m)
        input.Scanf("%d %d", &n, &m)
        InitQueue(&queue, n)
        list := make([]vertex, 0)
        for i := 0; i < n; i++ {
                var v vertex
                v.mark = false
	        v.num = i
	        v.edge = make([]int, 0)
                //list[i] = v
                list = append(list, v)
        }
        for i := 0; i < m; i++ {
                var u, v int 
                //fmt.Scan(&u, &v)
	        input.Scanf("%d %d", &u, &v)
                list[u].edge = append(list[u].edge, v)
              list[v].edge = append(list[v].edge, u)
        }
        //fmt.Scan(&k)
        input.Scanf("%d", &k) 
        dist := make([][]int, k)
        for i := 0; i < k; i++ {
                var a int
                //fmt.Scan(&a)
                input.Scanf("%d", &a)
                dist[i] = make([]int, n)
                BFS(a, &queue, list, &dist[i])
        }
        array := make([]int, 0)
        var ok bool
        for i := 0; i < n; i++ {
                ok = true
                for j := 0; j < k - 1; j++ {
                        if dist[j+1][i] != dist[j][i] || dist[j][i] == 0 {
                                ok = false;
	                	break;
                        }
                }
                if ok {
                        array = append(array, i)
                }
        }
        sort.Ints(array)
        if len(array) == 0 {
                fmt.Printf("-")
        } else {
                for i, _ := range array {
                        fmt.Printf("%d ", array[i])
        	}
        }
}