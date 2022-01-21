package main

import (
	"fmt"
	"os"
	"bufio"
//	"github.com/skorobogatov/input"
)

func main() {
        mindist, a, b := 1000000, 0, 0
       // var s, xstr, ystr string
        in := bufio.NewReader(os.Stdin)
        s, _ := in.ReadString('\n')
        str := []rune(s)
        x, _, _ := in.ReadRune()
        in.ReadRune()
        y, _, _ := in.ReadRune()
        //var l int
        //s := input.Gets()
        //input.Scanf("%s\n", &s)
       // l = len(s)
//	input.Scanf("%s %s\n", &xstr, &ystr)
	xbuf, ybuf := -1, -1
	//xarr := make([]int, l)
	//yarr := make([]int, l)
//	x, y := ([]rune)(xstr)[0], ([]rune)(ystr)[0] 
	i := 0
	for _, w := range str {        
	        if w == x {
	                xbuf = i
	                a+=1
	                if b != 0 {
	                        if xbuf - ybuf < mindist {
	                                mindist = xbuf - ybuf 
	                        }
	                }
	        }
	        if w == y {
	                ybuf = i
	                b+=1
	                if a != 0 {
	                        if ybuf - xbuf < mindist {
	                                mindist = ybuf - xbuf
	                        }
	                }
	        }
	        i += 1
	}
	fmt.Println(mindist - 1)
}