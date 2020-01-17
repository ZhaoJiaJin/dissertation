package solu

import(
    "testing"
    "gonum.org/v1/gonum/mat"
    //"conjugate/utils"
)

func TestVMatrix(t *testing.T){
    v := []float64{1,2,3,4,5,6}
    ma := mat.NewDense(3,2,v)
    vmx := &VMatrix{ma}
    vexpect := []float64{1,3,5,2,4,6}
    expvmx := mat.NewDense(6,1,vexpect)

    if !mat.Equal(vmx,expvmx){
        t.Fatal("not equal")
    }
    /*utils.PrintMatrix(ma)
    utils.PrintMatrix(vmx)
    utils.PrintMatrix(vmx.T())
    utils.PrintMatrix(expvmx)*/
}
