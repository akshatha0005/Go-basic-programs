package main
import "fmt"
import "github.com/Akshatha/Fibonacci/math"
func main() {
var i int
var n int
fmt.Println("enter number n to generate n fibonnaci series")
 fmt.Scan(&n)
 fmt.Println("The",n,"Fibonacci series")
 for i<=n {
 	fmt.Println(Fibonacci.Fib_gen(i))
 	i++ }
 }