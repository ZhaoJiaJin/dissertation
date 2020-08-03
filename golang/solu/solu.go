package solu

import(
    "gonum.org/v1/gonum/mat"
    "conjugate/utils"
    //"log"
    "runtime"
)

type IteSolu struct{
    A mat.Matrix
    T mat.Matrix
    m int
    n int
    N int
    y []float64
    srcx int
    srcy int
    threadNum int
    AtTA mat.Matrix
}

func NewIteSolu(a,t mat.Matrix,m,n,bign,lvl int,y []float64, threadNum int)*IteSolu{
    res := new(IteSolu)
    res.A = a
    res.T = t
    res.m,res.n = m,n
    res.N = bign
    res.y = y
    res.srcx,res.srcy = utils.MostSqure(res.N)
    res.threadNum = threadNum

    At := res.A.T()
    var AtTA mat.Dense
    AtTA.Product(At,res.T,res.A)
    res.AtTA = &AtTA
    return res
}


func (sl *IteSolu)DX(x []float64)[]float64{
    res := make([]float64, len(x))
    operator := &DXOpe{
        X:x,
        Res:res,
        ThreadNum:sl.threadNum,
        Rowx:sl.N,
        Colx:sl.n,
        Srcx:sl.srcx,
        Srcy:sl.srcy,
    }
    operator.Calculate()
    return operator.Res
}

func (sl *IteSolu)calQx(x *mat.VecDense)mat.Vector{
    res := x.RawVector().Data
    //res := x
    //X = mat.NewDense(sl.n, sl.N,raw).T()
    res = sl.DX(res)
    runtime.GC()
    res = sl.DX(res)
    runtime.GC()
    return mat.NewVecDense(len(res),res)
}

func (sl *IteSolu)calBtCBx(x *mat.VecDense)mat.Vector{
    Xt := mat.NewDense(sl.n,sl.N,x.RawVector().Data)
    X := Xt.T()
    res := new(mat.Dense)
    res.Product(X, sl.AtTA)
    return mat.DenseCopyOf(&VMatrix{res}).ColView(0)
}

func (sl *IteSolu)CalR(x *mat.VecDense)float64{
    b := sl.calb()
    runtime.GC()
    r := vectorsub(b, sl.calQx(x), sl.calBtCBx(x))
    return mat.Dot(r,r)
}

func (sl *IteSolu)FindSolution()(*mat.VecDense){
    b := sl.calb()
    runtime.GC()

    leng,_ := b.Dims()
    x := mat.NewVecDense(leng, nil)

    //x = make([]float64,leng)
    r := vectorsub(b, sl.calQx(x), sl.calBtCBx(x))
    runtime.GC()
    p := r

    r_k_norm := mat.Dot(r,r)
    //ori_norm := r_k_norm
    q := new(mat.VecDense)
    for i:=1; i<2*leng; i ++{
        //log.Println("Begin Itr:", i)
        q.Reset()
        q.AddVec(sl.calQx(p),sl.calBtCBx(p))
        runtime.GC()
        alpha := r_k_norm / mat.Dot(p,q)
        x.AddScaledVec(x,alpha,p)
        //if i % 5 == 0{
            r = vectorsub(b, sl.calQx(x), sl.calBtCBx(x))
            runtime.GC()
        //}else{
        //    r.AddScaledVec(r, alpha*-1, q)
        //}

        r_k1_norm := mat.Dot(r,r)
        beta := r_k1_norm/r_k_norm
        r_k_norm = r_k1_norm
        if r_k1_norm < 1e-5{
            //r = vectorsub(b, sl.calQx(x), sl.calBtCBx(x))
            runtime.GC()
            //log.Println("Itr:", i, r_k1_norm, mat.Dot(r,r),ori_norm)
            break
        }
        p.AddScaledVec(r,beta,p)

    }
    return x
}

func (sl *IteSolu)calb()mat.Vector{
    Yt := mat.NewDense(sl.m, sl.N, sl.y)
    Y := Yt.T()
    var res mat.Dense
    res.Product(Y, sl.T, sl.A)
    y := mat.DenseCopyOf(&VMatrix{&res})
    return y.ColView(0)
}
