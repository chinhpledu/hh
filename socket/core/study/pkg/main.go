package main

import "fmt"
import mym "myMath/math"

func main() {
    xs := []float64{1, 2, 3, 4}
    avg := mym.Average(xs)
    fmt.Println(avg)
}
