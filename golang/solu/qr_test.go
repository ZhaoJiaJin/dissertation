package solu

import(
    "testing"
    "gonum.org/v1/gonum/mat"
    "fmt"
)

func TestReducedQR(t *testing.T){
    inputarr := []float64{
        12,-51,
        6,167,
        -4,24,
    }
    A := mat.NewDense(3,2,inputarr)
    Q,R := ReducedQR(A)
    fmt.Printf("Q:\n%1.3f\n\n", mat.Formatted(Q))
    fmt.Printf("R:\n%1.3f\n\n", mat.Formatted(R))
}
