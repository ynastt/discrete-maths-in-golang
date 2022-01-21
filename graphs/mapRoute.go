//реализуем алгоритм Дейкстры для поиска кратчашего пути
//для этого понадобится очередь с приоритетами
package main

import(
	"fmt"
)

var n int
var inf = 8888888888

//priority queue
type queue struct {
        capc int
        heap []*vertex
        count int
}

var q queue

func InitPriorityQueue(n int) {
        q.heap = make([]*vertex, n)
        q.capc = n
        q.count = 0
}

func QueueEmpty(q *queue) bool {
  	return q.count == 0
}

func QueueInsert(q *queue, v *vertex) {
        i := q.count
        if i == q.capc {
                fmt.Println("error: переполнение")
        }
        q.count = i + 1
        q.heap[i] = v
        for i > 0 && q.heap[i].key < q.heap[(i - 1)/2].key {
                q.heap[(i - 1)/2], q.heap[i] = q.heap[i], q.heap[(i - 1)/2]
                q.heap[i].index = i
                i = (i - 1)/2
        }
        q.heap[i].index = i
        //(*v).index = i
}

func ExtractMin(q *queue) *vertex {
        if q.count == 0 {
                fmt.Println("error: очередь пустая")
        }
        ptr := q.heap[0]
        q.count -= 1
        if q.count > 0 {
                q.heap[0] = q.heap[q.count]
                q.heap[0].index = 0
                Heapify(0, q.count, q.heap)
        }
        return ptr
}

func Heapify (i int, n int, P []*vertex) {
	for {
		l := 2*i + 1
		r := l + 1
		j := i
		if l < n && P[i].key > P[l].key { 
                        i = l 
                }
		if r < n && P[i].key > P[r].key {
                        i = r 
                }
		if i == j {
                        break 
                }
		P[i], P[j] = P[j], P[i]
		P[i].index = i
		P[j].index = j
	}
}

func DecreaseKey(q *queue, i int, k int) {
	q.heap[i].key = k
	for i > 0 && q.heap[(i-1)/2].key > k {
		q.heap[(i-1)/2], q.heap[i] = q.heap[i], q.heap[(i-1)/2]
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	q.heap[i].index = i
}

type vertex struct {
	x int
  	y int
  	index int
  	weight int
  	key int //dist
  	edges []edge
}

type edge struct {//ребро как координаты инцидентной вершины
	x int
  	y int
}

func Relax(u *vertex, v *vertex, w int ) bool {
	//bool c
  	changed := u.key + w < v.key
  	if changed {
    		v.key = u.key + w 
  	}
 	return changed
}

func Dijkstra(g *[][]vertex) {
	InitPriorityQueue(n * n)
  	for i := 0; i < n; i++ {
    		for j := 0; j < n; j++ {
      			v := (*g)[i][j]
      			if j == 0 && i == 0 { //v==s, но s=(1,1), а счет идет с 0
        			v.key = v.weight //не 0!
      			} else {
        			v.key = inf
      			}
    			QueueInsert(&q, &v)
    		}  
  	}
 	for !QueueEmpty(&q) {
   		v := ExtractMin(&q)
   		v.index = -1
   		EG := v.edges
   		for _, el := range EG {
     			if (*g)[el.x][el.y].index != -1 &&
        		   Relax(v, &(*g)[el.x][el.y], (*g)[el.x][el.y].weight) {
				DecreaseKey(&q, (*g)[el.x][el.y].index, (*g)[el.x][el.y].key)
			}
   		}
 	}
}

func main() {
	var x int
        fmt.Scan(&n)
	gmap := make([][]vertex, n)
	for i := 0; i < n; i++ { 
      		gmap[i] = make([]vertex, n) 
    	}
  	for i := 0; i < n; i++ {
    		for j := 0; j < n; j++ {
			gmap[i][j].edges = make([]edge, 0)
		}
  	}
   	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
           		fmt.Scan(&x)
			gmap[i][j].weight = x
			gmap[i][j].x, gmap[i][j].y = i, j
            		var e edge
           		e.x, e.y = i, j
          		//разберёмся с ребрами
          		if i+1 < n {
            			gmap[i + 1][j].edges = append(gmap[i + 1][j].edges, e)
          		}
          		if j + 1 < n {
            			gmap[i][j + 1].edges = append(gmap[i][j + 1].edges, e)
          		}
          		if i > 0 {
            			gmap[i - 1][j].edges = append(gmap[i - 1][j].edges, e)
          		}
        		if j > 0 {
            			gmap[i][j - 1].edges = append(gmap[i][j - 1].edges, e)
        		}
		}
	}
  	Dijkstra(&gmap)
  	res := gmap[n-1][n-1].key
  	fmt.Println(res)
}

