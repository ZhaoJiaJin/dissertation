package solu

import(
    "gonum.org/v1/gonum/mat"
)

type VMatrix struct{
    mtx mat.Matrix
}

func (this *VMatrix)Dims()(int,int){
    r,c := this.mtx.Dims()
    return r*c,1
}

func (this *VMatrix)At(i, j int) float64{
    if j != 0{
        panic("j should be 0")
    }
    r,_ := this.mtx.Dims()
    return this.mtx.At(i%r,i/r)
}

func (this *VMatrix)T()mat.Matrix{
    return mat.Transpose{this}
}
