package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
            j, k, l = iota, iota, iota // j=9,k=9,l=9
            m = iota + 100
    )

    const z = iota // z=0

    fmt.Println(a,b,c,d,e,f,g,h,i,j,k,l,m,z)
}
