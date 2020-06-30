package solu

import(
    "gonum.org/v1/gonum/mat"
    "fmt"
    "log"
)

// getShat get matrix s_hat, s_hat = A^T  T A P^{-1}, the P matrix is identity matrix in our case, and sHat should be a symmetric matrix
func getShat(a,t mat.Matrix)(*mat.SymDense){
    ret := new(mat.Dense)
    ret.Product(a.T(),t,a)
    row,col := ret.Dims()
    if row != col{
        log.Panic("Matrix SHat should be symmetric, it should have equal rows and columns.")
    }
    return mat.NewSymDense(row, ret.RawMatrix().Data)
}


// validSHat check if we can get sHat from eigenvectors and eigenvalues
func validSHat(U *mat.Dense, egvals []float64){
    row,_ := U.Dims()
    ret := new(mat.Dense)
    R := mat.NewDiagDense(row,egvals)
    ret.Product(U,R,U.T())
    fmt.Printf("Recomputed S Hat:\n%1.3f\n\n", mat.Formatted(ret))
}


// getYHat computes yHat  YHat = YTAP−1 
func getYHat(oriy []float64,a,t mat.Matrix, m,n,bign int)(mat.Matrix){
    //rearrange y
    Y := mat.NewDense(bign,m,nil)
    for idx,_ := range oriy{
        col := idx/bign
        row := idx%bign
        Y.Set(row,col,oriy[idx])
    }

    yhat := new(mat.Dense)
    yhat.Product(Y,t,a)
    return yhat
}


// NewSylSolu calculate result using Sylvester method
func NewSylSolu(a,t mat.Matrix,m,n,bign,lvl int, y []float64, threadNum int)(ret mat.Vector){
    //Compute eigendecomposition SHat = URUT
    sHat := getShat(a,t)
    fmt.Printf("S Hat:\n%1.3f\n\n", mat.Formatted(sHat))
    //find eigenvectors and eigenvalues of sHat
    var eig mat.EigenSym
    ok := eig.Factorize(sHat, true)
    if !ok {
        log.Fatal("Eigendecomposition failed")
    }
    row,_ := sHat.Dims()
    egvals := make([]float64,row)
    eig.Values(egvals)
    var egvects mat.Dense
    eig.VectorsTo(&egvects)
    //validSHat(&egvects, egvals)
    yhat := getYHat(y,a,t,m,n,bign)
    fmt.Printf("Y Hat:\n%1.3f\n\n", mat.Formatted(yhat))
    return
}
