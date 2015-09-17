package main
import "fmt"
import "github.com/Akshatha/perimeter/math"
func main() {
	var l,b,r float64
    fmt.Println("enter the length and breath of rectangle")
    fmt.Scan(&l, &b)
    fmt.Println("enter the radius of circle")
    fmt.Scan(&r)
    fmt.Println("perimeter of rectangle is :", Perimeter.PerimeterR([]float64 {l,b}))
    fmt.Println()
    fmt.Println("perimeter of circle is :", Perimeter.PerimeterC(r))
    fmt.Println()
    }