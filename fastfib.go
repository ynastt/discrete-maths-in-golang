package main
import ("fmt"
        "math/big")

func rowcol(val1, val2, val3, val4 *big.Int) *big.Int {
       z := new(big.Int)
       y := new(big.Int)
       return z.Add(z.Mul(val1, val2),y.Mul(val3, val4) )
}

func main() {
        var n int
        var a, b, c, d *big.Int
        m1, m2, m3, m4 := big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(0)
        f1, f := big.NewInt(1), big.NewInt(0)
        fmt.Scanf("%d", &n)
        for n != 0 {
                if n%2 == 1 {
                        var x *big.Int
                        x = f
                        f = rowcol(f, m1, f1, m3)
                        f1 = rowcol(x, m2, f1, m4)
                }
                a, b, c, d = m1, m2, m3, m4
                /*aa+bc  ab+bd
                  ca+dc  cb+db*/
                m1 = rowcol(a, a, b, c)
                m2 = rowcol(a, b, b, d)
                m3 = rowcol(c, a, d, c)
                m4 = rowcol(c, b, d, d)
                n/=2
        }
        fmt.Println(f)
}