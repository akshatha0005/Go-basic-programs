package Fibonacci
func Fib_gen(x int) int {
	if x == 0 {
    return 0
    } else if x == 1 {
        return 1
    } else {
       return Fib_gen(x-1)+Fib_gen(x-2) }
   }