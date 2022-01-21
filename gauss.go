package main
import "fmt"

type frac struct {
	n, d int
}

func myabs(s int) int{
        if (s >= 0) {
                return s
        } else {
                return -s
        }
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return myabs(a)
}

func normal(f frac) frac {
	if f.n == 0 {
		f.d = 1
		return f
	}
        s := 1
        if (f.n > 0 && f.d < 0) || (f.n < 0 && f.d > 0){
                s = -1
        }
	x, y := myabs(f.n), myabs(f.d)
        z := gcd(x, y)
	f.n, f.d = s*x/z, y/z
	return f
}

func additionFrac(x, y frac) frac {
        var z frac
        z.n = x.n * y.d + y.n * x.d
        z.d = x.d * y.d
        return normal(z)
}

func subtractionFrac(x, y frac) frac {
        var z frac
        z.n = x.n * y.d - y.n * x.d
        z.d = x.d * y.d
        return normal(z)
}

func reversedFrac(f frac) frac{
        f.n, f.d = f.d, f.n
        return normal(f)
}

func negativeFrac(f frac) frac{
        f.n = -1 * f.n 
        return normal(f)
}

func multiply(a, b frac) frac {
        a.n *= b.n
        a.d *= b.d
        return normal(a)
}

func swapRows(m [][]frac, row1 int , row2 int, n int) {
        for j := 0; j < n+1; j++ {
		m[row1][j], m[row2][j] = m[row2][j], m[row1][j]
	}
}

func subtractRows(m [][]frac, k frac, row1, row2, n int) { 
	for j := 0; j < n+1; j++ {
		m[row1][j] = subtractionFrac(m[row1][j], multiply(m[row2][j],k))
	}
}
//приведение к ступенчатому виду (верхнетреугольному)
func matrixToSteppedView(matrix [][]frac, n int) {
        for i :=0; i < n; i++ {
          //fmt.Println("it`s ", i)
                m := matrix[i][i]
                //ноль не должен быть на главной диагонали
                for k := i + 1; k < n && matrix[i][i].n == 0; k++ { 
                        swapRows(matrix, i, k, n)
                }
                m = matrix[i][i]
     // fmt.Println()
       // printMatrix(matrix, n)
                //сделаем элемент в углу минора 1/1
                for j := 0; j < n+1; j++ { 
                        matrix[i][j] = multiply(matrix[i][j],reversedFrac(m))
                }
        //fmt.Println()
        //printMatrix(matrix, n)
                for c := i + 1; c < n; c++ {
			if matrix[c][i].n == 0 { //уже зануленный элеменет => пропускаем
		                continue
	                } 
                q := matrix[c][i]
                subtractRows(matrix, q, c, i, n)
		}
        //fmt.Println()
       // printMatrix(matrix, n)
	}
//printMatrix(matrix, n)
}

//решение слау
func SLE(matrix [][]frac, n int) (bool, []frac) {
	ok := true
	result := make([]frac, n)
	for i := n - 1; i >= 0; i-- {
		if  matrix[i][i].n == 0 {
			ok = false
		}
		x := matrix[i][n]
		for j := i + 1; j < n; j++ {
                        xn := multiply(negativeFrac(matrix[i][j]),result[j])
			x = additionFrac(x, xn)
		}
		result[i] =x
	}
	return ok, result
}


func printMatrix(a [][]frac, n int) {
	fmt.Println()
	for i := 0; i < n; i++ {
		for j := 0; j < n+1; j++ {
			fmt.Print(a[i][j].n, "/", a[i][j].d, " ")
		}
		fmt.Println()
	}
}
func main() {
        var n int
        fmt.Scan(&n)
        var matrix [][]frac = make([][]frac, n+1)
        for i := 0; i < n; i++ {
                matrix[i] = make([]frac, n + 1)
        }
        for i := 0; i < n; i++ {
	        for j := 0; j <= n; j++ {
 		        fmt.Scan(&matrix[i][j].n)
                        matrix[i][j].d = 1
                }
        }
        matrixToSteppedView(matrix, n)
        //вывод значений х  
        ok, res := SLE(matrix, n)
	if ok {
                for _, x := range res {
			fmt.Printf("%d/%d\n", x.n, x.d)
		}	
	} else {
		fmt.Printf("No solution")
	}
}