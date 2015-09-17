package Perimeter
import "testing"
type testpair struct {
  n  float64 
  result float64 
}
type testpair1 struct {
  values []float64
  result1 float64
}

var tests1 = []testpair1 {
  { []float64{0,2}, 0 },
  { []float64{2,3 }, 10 },
  { []float64{20,30}, 100 },
}

var tests = []testpair{
  {  0, 0 },
  { 5,  31.41592653589793 },
}

func Testperi(t *testing.T) {

  for _, pair := range tests {
    v := PerimeterC(pair.n)
    if v != pair.result {
      t.Error(
        "For", pair.n,
        "expected", pair.result,
        "got", v,
      )
    }
  }
}

func Testperi1(t *testing.T) {
  for _, pair := range tests1 {
    v := PerimeterR(pair.values)
    if v != pair.result1 {
      t.Error(
        "For", pair.values,
        "expected", pair.result1,
        "got", v,
      )
    }
  }
}
