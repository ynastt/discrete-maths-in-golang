package main
import "fmt"

type vertex struct {
        num int
        index int
        key int
        value int
        edges []edge
}

type edge struct {
        prevVertex int //u
        nextVertex int //v
        l int //len
}
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

func Insert(q *queue, v *vertex) {
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

func MST_Prim(graph []*vertex, n int) int {
        min := 0
        for i, _ := range graph {
                graph[i].index = -1
        }
        InitPriorityQueue(n - 1)
        v := graph[0]
        for {
                v.index = -2
                EG := v.edges
                for _, e := range EG {
                        u := e.nextVertex
                        a := e.l
                        w := e.prevVertex
                        if graph[u].index == -1 {
                                graph[u].key = a
                                graph[u].value = w
		                Insert(&q, graph[u])
                        } else {
                                if graph[u].index != -2 && a < graph[u].key {
                                        graph[u].value = w
                                        DecreaseKey(&q, graph[u].index, a)
                                }
                        }
                }
                if q.count == 0 {
                        break 
                }
	        v = ExtractMin(&q)
	        min += v.key
        }
        return min
} 

func main() {
        var n, m, u, v, l int
        fmt.Scan(&n, &m)
        list := make([]*vertex, 0)
        for i := 0; i < n; i++ {
                var v vertex
	        v.key = -1
	        v.edges = make([]edge, 0)
                list = append(list, &v)
        }
        for i := 0; i < m; i++ {
                fmt.Scan(&u, &v, &l)
	        var e edge
                e.prevVertex = u
                e.nextVertex = v
                e.l = l
                list[u].edges = append(list[u].edges, e)
                e.prevVertex = v
                e.nextVertex = u
                list[v].edges = append(list[v].edges, e) 	
	}     
        min := MST_Prim(list, n)    
        fmt.Println(min)
}