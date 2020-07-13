package solu


import(
    "gonum.org/v1/gonum/mat"
    "math"
)

//ReducedQR perform reduced QR decomposition, A is m*n matrix, m>=n, and Q is m*n matrix, R is n*n matrix
func ReducedQR(A *mat.Dense)(Q *mat.Dense, R *mat.Dense){
    m,n := A.Dims()
    if m < n{
        panic("wrong matrix size")
    }
    unorms := make([]float64,0,n)
    unormsqrt := make([]float64,0,n)
    uvectors := make([]mat.Vector,0,n)
    for cur := 0; cur < n; cur ++{
        oria := A.ColView(cur)
        cura := mat.VecDenseCopyOf(oria)
        for pre := 0; pre < cur; pre ++{
            coffi := VectorMul(uvectors[pre],oria) / unormsqrt[pre]
            cura.AddScaledVec(cura,-1 * coffi,uvectors[pre])
        }
        uvectors = append(uvectors, cura)
        nmsqur := VectorNorm(cura)
        unormsqrt = append(unormsqrt,nmsqur)
        unorms = append(unorms,math.Sqrt(nmsqur))
    }

    Q = mat.NewDense(m,n,nil)
    for cidx := 0; cidx < n; cidx ++{
        for ridx := 0; ridx < m; ridx ++{
            Q.Set(ridx,cidx,uvectors[cidx].AtVec(ridx) / unorms[cidx])
        }
    }
    R = new(mat.Dense)
    R.Product(Q.T(), A)
    return Q,R
}

func VectorNorm(a mat.Vector)(float64){
    return VectorMul(a,a)
}

func VectorMul(a mat.Vector, b mat.Vector)(float64){
    if a.Len() != b.Len(){
        panic("wrong vector length")
    }
    res := 0.0
    for i:=0; i < a.Len(); i ++{
        res += (a.AtVec(i)*b.AtVec(i))
    }
    return res
}
