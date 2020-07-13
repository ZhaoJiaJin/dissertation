package solu

import(
    "gonum.org/v1/gonum/mat"
    "conjugate/utils"
    "math"
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
    //find eigenvectors and eigenvalues of sHat
    srcx,srcy := utils.MostSqure(bign)
    R,U:= EigenSymFac(sHat)
    //validSHat(U, R)
    yhat := getYHat(y,a,t,m,n,bign)

    //We need to store Q B and H

    Qs := make([]*mat.Dense,0)
    Bs := make([]*mat.Dense,0)
    Hs := make([]*mat.Dense,0)
    V,B0 := ReducedQR(yhat)
    //fmt.Printf("y hat :\n%1.3f\n\n", mat.Formatted(yhat))
    // have V and B0 now
    //fmt.Printf("!!!!Y:\n%1.3f\n\n", mat.Formatted(yhat))
    //Bs = append(Bs, B0)
    r := mat.Norm(B0,2)
    i := 0
    var Vold *mat.Dense
    var Phi []float64
    var Qi *mat.Dense
    //var fullV *mat.Dense
    //fullV = addv(fullV, V)
    //validateT(fullV,lvl)
    B0norm := mat.Norm(B0,2)
    for math.Sqrt(r) > epsilon * B0norm{
        fmt.Println("Ite:",i)
        i ++
        // Lanczos steps
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
            VoldB.Product(Vold,Bs[len(Bs)-1])
            W.Sub(W,&VoldB)
        }
        H := new(mat.Dense)
        H.Product(V.T(),W)
        Hs = append(Hs,H)
        VH := new(mat.Dense)
        VH.Product(V,H)
        W.Sub(W,VH)

        Vold = V
        curB := new(mat.Dense)
        //fmt.Printf("W:\n%1.3f\n\n", mat.Formatted(W))
        V,curB = ReducedQR(W)
        //fmt.Printf("New V:\n%1.3f\n\n", mat.Formatted(V))
        //fmt.Printf("New V:\n%1.3f\n\n", mat.Formatted(V))
        Bs = append(Bs, curB)
        //fullV = addv(fullV,V)
        //fmt.Printf("New Full V:\n%1.3f\n\n", mat.Formatted(fullV))

        /* computing residual norm ******************************* */
        // combine H and B to get Ti
        T := mat.NewSymDense(n*i,nil)
        // Set Hi in the diagonal
        for hidx:=0; hidx < i; hidx ++{
            curh := Hs[hidx]
            for ridx := 0; ridx < n; ridx ++{
                for cidx := ridx; cidx < n; cidx ++{
                    T.SetSym(hidx*n + ridx, hidx * n + cidx, curh.At(ridx,cidx))
                }
            }
        }
        // Set Bi in the upper diagonal
        for bidx := 0; bidx < i-1; bidx ++{
            curb := Bs[bidx]
            for ridx := 0; ridx < n; ridx ++{
                for cidx := 0; cidx < n; cidx ++{
                    T.SetSym(bidx*n+ridx,bidx*n+n+cidx, curb.At(cidx,ridx))
                }
            }
        }
        //fmt.Printf("T :\n%1.3f\n\n", mat.Formatted(T))
        //validateT(fullV,lvl)
        Phi,Qi = EigenSymFac(T)
        Qs = append(Qs,Qi)
        E1 := GetEi(0, i, n)
        var UTB0T mat.Dense
        UTB0T.Product(U.T(),B0.T())
        var E1TQi mat.Dense
        E1TQi.Product(E1.T(), Qi)
        var S mat.Dense
        S.Product(&UTB0T,&E1TQi)
        Em := GetEi(i-1,i,n)
        var J mat.Dense
        J.Product(Qi,Em, curB.T())
        //fmt.Printf("T:\n%1.3f\n\n", mat.Formatted(&S))
        //fmt.Printf("T:\n%1.3f\n\n", mat.Formatted(&J))
        rowS,_ := S.Dims()
        r = 0
        for j := 0; j < n; j ++{
            Kdiagonal := make([]float64,len(Phi))
            for tmpidx := 0; tmpidx < len(Phi); tmpidx ++{
                Kdiagonal[tmpidx] = 1/(Phi[tmpidx] + R[j])
            }
            Kinverse := mat.NewDiagDense(len(Kdiagonal),Kdiagonal)
            ej := make([]float64,rowS)
            ej[j] = 1
            ejmat := mat.NewDense(rowS,1,ej)
            var deltaVec mat.Dense
            deltaVec.Product(ejmat.T(),&S,Kinverse,&J)
            //fmt.Printf("delta:\n%1.3f\n\n", mat.Formatted(&deltaVec))
            //r += 
            var deltaV mat.Dense
            deltaV.Product(&deltaVec,deltaVec.T())
            //fmt.Printf("delta:\n%1.3f\n\n", mat.Formatted(&deltaV))
            rowdelta,coldelta := deltaV.Dims()
            if rowdelta != 1 || coldelta != 1{
                panic("residual delta should have wrong size")
            }
            r += deltaV.At(0,0)
        }
        fmt.Println("new residual",r)

    }
    E1 := GetEi(0, i, n)
    var QE1B0U mat.Dense
    QE1B0U.Product(Qi.T(),E1,B0,U)
    rows,cols := QE1B0U.Dims()
    ek := mat.NewDense(1,rows,nil)
    er := mat.NewDense(cols,1,nil)
    F := mat.NewDense(i*n,n,nil)
    for k:=0; k < i*n; k++{
        for rho:=0; rho < n; rho++{
            ek.Set(0,k,1)
            er.Set(rho,0,1)
            var curF mat.Dense
            curF.Product(ek,&QE1B0U,er)
            F.Set(k,rho,curF.At(0,0)/(Phi[k] + R[rho]))
            ek.Set(0,k,0)
            er.Set(rho,0,0)
        }
    }
    var Z mat.Dense
    Z.Product(Qi,F,U.T())
    M := mat.NewDense(bign,n,nil)
    B0Inver := new(mat.Dense)
    B0Inver.Inverse(B0)
    V = new(mat.Dense)
    V.Product(yhat,B0Inver)
    Zrawdata := Z.RawMatrix().Data
    //fmt.Printf("Z:\n%1.3f\n\n", mat.Formatted(&Z))
    for j:=1; j<=i; j++{
        Gi := mat.NewDense(n,n,Zrawdata[n*n*(j-1):n*n*j])
        //fmt.Printf("G:\n%1.3f\n\n", mat.Formatted(Gi))
        var VGi mat.Dense
        VGi.Product(V,Gi)
        M.Add(M,&VGi)

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
        if j>1{
            var VoldB mat.Dense
            VoldB.Product(Vold,Bs[j-2])
            W.Sub(W,&VoldB)
        }
        VH := new(mat.Dense)
        VH.Product(V,Hs[j-1])
        W.Sub(W,VH)
        Vold = V
        BInverse := new(mat.Dense)
        BInverse.Inverse(Bs[j-1])
        V.Product(W,BInverse)
    }
    resdata := make([]float64,0,bign*n)
    for colidx := 0; colidx < n; colidx ++{
        for rowidx := 0; rowidx < bign; rowidx ++{
            resdata = append(resdata, M.At(rowidx,colidx))
        }
    }
    res := mat.NewVecDense(bign*n, resdata)
    return res
}


func QRFac(A mat.Matrix)(*mat.Dense,*mat.Dense){
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

func EigenSymFac(A mat.Symmetric)([]float64, *mat.Dense){
    var eig mat.EigenSym
    ok := eig.Factorize(A, true)
    if !ok {
        log.Fatal("Eigendecomposition failed")
    }
    row,_ := A.Dims()
    egvals := make([]float64,row)
    eig.Values(egvals)
    var egvects mat.Dense
    eig.VectorsTo(&egvects)
    return egvals, &egvects
}

func GetEi(idx, j, n int)(*mat.Dense){
    res := mat.NewDense(j*n,n,nil)
    startrow := idx * n
    for curidx := 0; curidx < n; curidx ++{
        res.Set(startrow + curidx, curidx, 1)
    }
    return res
}

func addv(full *mat.Dense,v mat.Matrix)(*mat.Dense){
    if full == nil{
        return mat.DenseCopyOf(v)
    }
    orow,ocol := full.Dims()
    vrow,vcol := v.Dims()
    ret := mat.NewDense(orow,ocol+vcol,nil)
    for i:=0; i < orow; i++{
        for j:=0; j < ocol; j++{
            ret.Set(i,j,full.At(i,j))
        }
    }
    for i:=0; i < vrow; i++{
        for j:=0; j < vcol; j++{
            ret.Set(i,ocol+j,v.At(i,j))
        }
    }
    return ret
}

func validateT(V *mat.Dense, lvl int){
    d := generateD(lvl)
    var ttt mat.Dense
    ttt.Product(V.T(),d,d,V)
    fmt.Printf("TV:\n%1.3f\n\n", mat.Formatted(&ttt))
}
