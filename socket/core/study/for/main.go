package main

import "fmt"

func main(){
  var i int = 1
  for i <= 10 {
      fmt.Println(i)
      i++
  }

  for i:= 1; i <=10; i++{
    fmt.Println("Phan tu thu " , i)
    if (i % 2 == 0){
       fmt.Println("Chan")
    }
    switch i {
    case 1: fmt.Println("Mot")
    case 2: fmt.Println("Hai")
    case 3: fmt.Println("Ba")
    default: fmt.Println("Khong biet")
    }
  }

  var a [5]int
  slice1 := []int{1, 2, 3}
  a[4] = 1000
  a[2] = 22

  var total float64 = 0

  for index := 0 ; index < len(a); index++ {
      a[index] = index  + 1
      total = total + float64(a[index])
  }
  fmt.Println(a, total)

  var all float64 = 1

  for _, value := range a {
      all = all * float64(value)
  }
  slice2 := make([]int, 2)
  copy(slice2, slice1)
  fmt.Println(slice1,slice2, all)

  y := make(map[string]int)
  y["test1"] = 1
  y["test2"] = 2

  value, ok := y["test3"]
  fmt.Println(y,value,ok)
}
