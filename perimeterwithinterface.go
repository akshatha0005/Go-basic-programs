package main
import "fmt"
import "math"
type shape interface
 { 
 	perimeter() float64
 }
type rectangle struct 
{ 
	length, width float64
	 }
type circle struct 
{ 
	radius float64 
}
func (p rectangle) perimeter() float64{ 
	return 2*p.length + 2*p.width
	 } 
func (q circle) perimeter() float64{ 
	return 2 * math.Pi * q.radius
	 }
func main() {
	var l,b,r float64
    var s shape 
    fmt.Println("enter the length and breath of rectangle")
    fmt.Scan(&l, &b)
    fmt.Println("enter the radius of circle")
    fmt.Scan(&r)
    p := rectangle{l,b}
    s = p
    fmt.Println("perimeter of rectangle is :", s.perimeter())
    fmt.Println()
    q:=circle{r} 
    s = q
    fmt.Println("perimeter of circle is :", s.perimeter())
    fmt.Println()
    }