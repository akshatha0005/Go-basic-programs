package Perimeter
import "math"
func PerimeterC(r float64) float64{ 
	return 2 * math.Pi * r
	 } 
func PerimeterR(xs []float64) float64 {
  result := 2.0*(xs[0]+xs[1])
  return result
}