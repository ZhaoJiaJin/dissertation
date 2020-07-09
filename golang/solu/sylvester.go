package solu

import(
    "gonum.org/v1/gonum/mat"
    "conjugate/utils"
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
func NewSylSolu(a,t mat.Matrix,m,n,bign,lvl int, y []float64, threadNum int, epsilon float64)(ret mat.Vector){
    //Compute eigendecomposition SHat = URUT
    sHat := getShat(a,t)
    fmt.Printf("S Hat:\n%1.3f\n\n", mat.Formatted(sHat))
    //find eigenvectors and eigenvalues of sHat
    var eig mat.EigenSym
    ok := eig.Factorize(sHat, true)
    if !ok {
        log.Fatal("Eigendecomposition failed")
    }
    srcx,srcy := utils.MostSqure(bign)
    row,_ := sHat.Dims()
    egvals := make([]float64,row)
    eig.Values(egvals)
    var egvects mat.Dense
    eig.VectorsTo(&egvects)
    //validSHat(&egvects, egvals)
    yhat := getYHat(y,a,t,m,n,bign)
    fmt.Println(yhat.Dims())

    //We need to store Q B and H

    //Qs := make([]*mat.Dense,0)
    Bs := make([]*mat.Dense,0)
    Hs := make([]*mat.Dense,0)
    V,B0 := FacQR(yhat)
    //fmt.Printf("y hat :\n%1.3f\n\n", mat.Formatted(yhat))
    // have V and B0 now
    Bs = append(Bs, B0)
    r := mat.Norm(B0,2)
    i := 0
    var Vold *mat.Dense
    for r > epsilon * mat.Norm(B0,2){
        i ++
        // Lanczos steps
        // cal W = AV
        // we do it column by column
        W := mat.NewDense(bign,n,nil)
        for colidx:=0; colidx < n; colidx ++{
            coli := make([]float64,bign)
            for r := 0; r < bign; r ++{
                coli[r] = V.At(r,colidx)
            }
            d1res := make([]float64, bign)
            operator := &DXOpe{
                X:coli,
                Res:d1res,
                ThreadNum:threadNum,
                Rowx:bign,
                Colx:1,
                Srcx:srcx,
                Srcy:srcy,
            }
            operator.Calculate()
            d2res := make([]float64, bign)
            operator = &DXOpe{
                X:d1res,
                Res:d2res,
                ThreadNum:threadNum,
                Rowx:bign,
                Colx:1,
                Srcx:srcx,
                Srcy:srcy,
            }
            operator.Calculate()

            W.SetCol(colidx,d2res)
        }
        /*d := generateD(lvl)

        //fmt.Printf("d:\n%1.1f\n\n", mat.Formatted(d))
        x1,x2 := d.Dims()
        if x1 != x2{
            fmt.Println("len",x1,x2)
            panic("not sym")
        }
        for i1 := 0; i1 < x1; i1 ++{
            for i2:=0; i2 < x2; i2++{
                if d.At(i1,i2) != d.At(i2,i1){
                    fmt.Println("pos",i1,i2,d.At(i1,i2),d.At(i2,i1))
                    panic("not sym")
                }
            }
        }
        d2 := new(mat.Dense)
        d2.Product(d,d)
        W.Product(d2,V)*/

        if i > 1{
            var VoldB mat.Dense
            VoldB.Product(Vold,Bs[i-1])
            W.Sub(W,&VoldB)
        }
        H := new(mat.Dense)
        H.Product(V.T(),W)
        fmt.Printf("H:\n%1.3f\n\n", mat.Formatted(H))
        Hs = append(Hs,H)
        VH := new(mat.Dense)
        VH.Product(V,H)
        W.Sub(W,VH)

        Vold = V
        curB := new(mat.Dense)
        V,curB = FacQR(W)
        Bs = append(Bs, curB)

        /* computing residual norm ******************************* */
        r = 0
        // combine H and B to get Ti
        T = mat.NewSymDense(n*i,nil)
        // Set Hi in the diagonal
        for hidx:=0; hidx < i; hidx ++{
            curh := Hs[hidx]
            for ridx := 0; ridx < n; ridx ++{
                for cidx = ridx; cidx < n; cidx ++{
                    T.SetSym(hidx*n + ridx, hidx * n + cidx, curh.At(ridx,cidx))
                }
            }
        }
        // Set Bi in the upper diagonal
        for bidx := 0; bidx < i-1; bidx ++{
            curb := Bs[bidx]
        }

    }
    return nil
}


func FacQR(A mat.Matrix)(*mat.Dense,*mat.Dense){
    bign,n := A.Dims()
    var qrfac mat.QR
    qrfac.Factorize(A)
    var fullV mat.Dense
    var fullB mat.Dense
    qrfac.QTo(&fullV)
    qrfac.RTo(&fullB)

    V := mat.NewDense(bign,n,nil)
    for r:=0; r < bign; r ++{
        for c := 0; c < n; c ++{
            V.Set(r,c,fullV.At(r,c))
        }
    }
    B := mat.NewDense(n,n,nil)
    for r:=0; r < n; r ++{
        for c := 0; c < n; c ++{
            B.Set(r,c,fullB.At(r,c))
        }
    }
    return V,B
}
