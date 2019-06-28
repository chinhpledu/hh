package main

import "fmt"

func zero(xPtr *int) {
    *xPtr = 0
}

func main() {
    x := 5
    zero(&x)
    fmt.Println(x)
    fmt.Println(&x)
    fmt.Println(f(1,2,45,33,56,566,5))
}
func f(args ...int) int{
   var total int = 0
   for i := 0 ; i < len(args) ; i ++ {
       total = total + args[i]
   }
   return total
}
