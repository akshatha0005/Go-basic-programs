package Fibonacci
import "testing"
	type testpair struct {
  n int 
  fib int 
}

var tests = []testpair{
  {  0, 0 },
  { 1, 1 },
  { 25, 75025} ,
  {  10, 55 } ,
}

func Testfib(t *testing.T) {
  for _, pair := range tests {
    v := Fib_gen(pair.n)
    if v != pair.fib {
      t.Error(
        "For", pair.n,
        "expected", pair.fib,
        "got", v,
      )
    }
  }
}