package main

import ("fmt"
        "github.com/skorobogatov/input"
        )

var array []vertex //массив всех вершин графа
var markOfComp int // какая компонента
var curComp, maxComp component //очередная и максимальная компоненты

type component struct {
	kol1 int //вершины
	kol2 int //ребра
	mark int //
}

type vertex struct {
	mark int
	visited bool
	kin []int //вершины, с которыми данная вершина имеет ребра
}

func updateComp(i int) {
	array[i].visited = true
	array[i].mark = markOfComp
	curComp.kol1 += 1
	curComp.kol2 += len(array[i].kin)
	for _, v := range array[i].kin {
		if !array[v].visited {
			updateComp(v)
		}
	}
}

func main() {
	var n, m int
	//fmt.Scan(&n, &m)
	input.Scanf("%d %d", &n, &m)
	array = make([]vertex, n)
	for _, v := range array {
		v.kin = make([]int, 0)
	}
	//вводим граф и
	//для каждой вершины в массив ее "родни" добавляем вершины т.е. 
	//вершины, с которыми она имеет ребро
	for i := 0; i < m; i++ {
		var u, v int
		input.Scanf("%d %d", &u, &v)
		//fmt.Scan(&u, &v)
		array[u].kin = append(array[u].kin, v)
		if u != v {
			array[v].kin = append(array[v].kin, u)
		}
	}
	markOfComp = 1
	for i, _ := range array {
		if !array[i].visited {
			curComp.kol1 = 0
			curComp.kol2 = 0
			curComp.mark = markOfComp
			updateComp(i)
		}
		if ((curComp.kol2 > maxComp.kol2 &&
		        curComp.kol1 == maxComp.kol1) ||
		        curComp.kol1 > maxComp.kol1) {
		                maxComp.kol1 = curComp.kol1;
			        maxComp.kol2 = curComp.kol2;
			        maxComp.mark = curComp.mark;
		}
		markOfComp += 1
	}
	fmt.Println("graph {")
	//вершины и ребра + пометка [color = red]
	for i, _ := range array {
		if array[i].mark  == maxComp.mark {
			fmt.Println(i, "[color = red]")
		} else {
			fmt.Println(i)
		}
	}
	for i := 0; i < n; i++ {
		for _, v := range array[i].kin {
          if i <= v{
          if  array[i].mark == maxComp.mark {
					fmt.Println(i, "--", v, "[color = red]")
		  } else {
            fmt.Println(i, "--", v)
          }
			
		}
        }
	}
	fmt.Println("}")
}