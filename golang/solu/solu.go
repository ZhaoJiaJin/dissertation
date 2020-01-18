package solu

import(
    "gonum.org/v1/gonum/mat"
    "conjugate/utils"
    "log"
    "runtime"
    "sync"
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
    AllNeigh [][]int
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
    res.AllNeigh = findAllNeigh(res.srcx,res.srcy, res.N,res.threadNum)
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
        AllNeigh:sl.AllNeigh,
    }
    operator.Calculate()
    return operator.Res
}

func (sl *IteSolu)calQx(x *mat.VecDense)mat.Vector{
    log.Println("begin Qx")
    res := x.RawVector().Data
    //res := x
    //X = mat.NewDense(sl.n, sl.N,raw).T()
    res = sl.DX(res)
    runtime.GC()
    res = sl.DX(res)
    runtime.GC()
    ret := mat.NewVecDense(len(res),res)
    log.Println("end Qx")
    return ret
}

func (sl *IteSolu)calBtCBx(x *mat.VecDense)mat.Vector{
    log.Println("begin BtCBx")
    Xt := mat.NewDense(sl.n,sl.N,x.RawVector().Data)
    X := Xt.T()
    res := new(mat.Dense)
    res.Product(X, sl.AtTA)
    ret := mat.DenseCopyOf(&VMatrix{res}).ColView(0)
    log.Println("end BtCBx")
    return ret
}

func (sl *IteSolu)calAx(x *mat.VecDense)(v1 mat.Vector,v2 mat.Vector){
    var wg sync.WaitGroup
    wg.Add(1)
    go func(){
        v1 = sl.calQx(x)
        wg.Done()
    }()
    wg.Add(1)
    go func(){
        v2 = sl.calBtCBx(x)
        wg.Done()
    }()
    wg.Wait()
    return
}

func (sl *IteSolu)FindSolution()mat.Vector{
    b := sl.calb()
    runtime.GC()

    leng,_ := b.Dims()
    x := mat.NewVecDense(leng, nil)

    //x = make([]float64,leng)
    v1,v2 := sl.calAx(x)
    r := vectorsub(b, v1,v2)
    runtime.GC()
    p := r

    r_k_norm := mat.Dot(r,r)
    //ori_norm := r_k_norm
    q := new(mat.VecDense)
    for i:=1; i<2*leng; i ++{
        log.Println("Begin Itr:", i)
        q.Reset()
        q.AddVec(sl.calAx(p))
        runtime.GC()
        alpha := r_k_norm / mat.Dot(p,q)
        x.AddScaledVec(x,alpha,p)
        //if i % 1 == 0{
            v1,v2 = sl.calAx(x)
            r = vectorsub(b, v1,v2)
            runtime.GC()
        //}else{
        //    r.AddScaledVec(r, alpha*-1, q)
        //}

        r_k1_norm := mat.Dot(r,r)
        log.Println("New R:",r_k1_norm)
        beta := r_k1_norm/r_k_norm
        r_k_norm = r_k1_norm
        if r_k1_norm < 1e-3{
            //r = vectorsub(b, sl.calQx(x), sl.calBtCBx(x))
            runtime.GC()
            //log.Println("Itr:", i, r_k1_norm, mat.Dot(r,r))
            log.Println("Itr:", i, r_k1_norm)
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
    //TODO: pool sl.y
    return y.ColView(0)
}
