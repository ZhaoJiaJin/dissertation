package solu

import(
    "gonum.org/v1/gonum/mat"
    "conjugate/utils"
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
}

func NewIteSolu(a,t mat.Matrix,m,n,bign int,y []float64, threadNum int)*IteSolu{
    res := new(IteSolu)
    res.A = a
    res.T = t
    res.m,res.n = m,n
    res.N = bign
    res.y = y
    res.srcx,res.srcy = utils.MostSqure(res.N)
    res.threadNum = threadNum
    return res
}


func (this *IteSolu)DX(x []float64)[]float64{
    res := make([]float64, len(x))
    operator := &DXOpe{
        X:x,
        Res:res,
        ThreadNum:this.threadNum,
        Rowx:this.N,
        Colx:this.n,
        Srcx:this.srcx,
        Srcy:this.srcy,
    }
    operator.Calculate()
    return operator.Res
}

func (this *IteSolu)getQx(x []float64)mat.Matrix{
    //res := x.RawMatrix().Data
    res := x
    //X = mat.NewDense(this.n, this.N,raw).T()
    res = this.DX(res)
    res = this.DX(res)
    return mat.NewDense(len(res),1,res)
}

func (this *IteSolu)FindSolution()[]float64{
    At := this.A.T()
    var AtTA mat.Dense
    AtTA.Product(At,this.T,this.A)

    b := this.calb()

    leng,_ := b.Dims()
    //_ = mat.NewDense(leng, 1, nil)
    _ = make([]float64,leng)

    //r = b - 

    return []float64{}
}

func (this *IteSolu)calb()mat.Matrix{
    Yt := mat.NewDense(this.m, this.N, this.y)
    Y := Yt.T()
    var res mat.Dense
    res.Product(Y, this.T, this.A)
    return &VMatrix{&res}
}
