package main

import ("fmt"
        "bytes")

func main() {
        var buf []byte
        var res = 0 
        fmt.Scan(&buf)
        for {
                k := bytes.Index(buf, []byte(")"))
                if k == -1{
                        break
                }
                res++        
        //Для правильности работы программы необходимо,
        //чтобы в каждой скобке после удаления найденного выражения
        //число входящих выражений не изменилось.
        //=> делаем замену на пробел " "
                buf = bytes.Replace(buf, buf[k-4 : k+1], []byte(" "), -1) 
        }
        fmt.Println(res)
}