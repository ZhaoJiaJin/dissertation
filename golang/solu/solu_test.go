package solu

import(
    "testing"
    "fmt"
    "math/rand"
    "gonum.org/v1/gonum/mat"
    "conjugate/utils"
)

var debug = false

func randv()(float64){
    return float64(rand.Intn(50))
}

func TestCalb(t *testing.T){
    m := 9
    n := 4
    N := 4*12
    avl := make([]float64, m*n)
    for i := range avl{
        avl[i] = randv()
    }
    tvl := make([]float64, m * m)
    for i:=0; i < m; i++{
        tvl[i*m+i] = randv()
    }
    yvl := make([]float64, m*N)
    for i := range yvl{
        yvl[i] = randv()
    }
    ma := mat.NewDense(m,n,avl)
    mt := mat.NewDense(m,m,tvl)
    my := mat.NewDense(m*N,1,yvl)
    sl := NewIteSolu(ma,mt,m,n,N,yvl,5)
    b := sl.calb()
    if debug{
        fmt.Println("b")
        utils.PrintMatrix(b)
    }

    IN := GenerateI(N)
    B := new(mat.Dense)
    B.Kronecker(ma, IN)
    C := new(mat.Dense)
    C.Kronecker(mt,IN)
    expect := new(mat.Dense)
    expect.Product(B.T(), C, my)
    if !mat.Equal(expect, b){
        t.Fatal("getb wrong")
    }
}

func TestQx(t *testing.T){
    m := 9
    n := 4
    lvl := 1
    N := utils.CalN(lvl)

    avl := make([]float64, m*n)
    for i := range avl{
        avl[i] = randv()
    }
    tvl := make([]float64, m * m)
    for i:=0; i < m; i++{
        tvl[i*m+i] = randv()
    }
    yvl := make([]float64, m*N)
    for i := range yvl{
        yvl[i] = randv()
    }
    ma := mat.NewDense(m,n,avl)
    mt := mat.NewDense(m,m,tvl)

    xvl := make([]float64, n*N)
    for i := range xvl{
        xvl[i] = randv()
    }
    x := mat.NewVecDense(n*N,xvl)
    sl := NewIteSolu(ma,mt,m,n,N,yvl,3)
    res := sl.calQx(x)

    d := generateD(lvl)
    d2 := new(mat.Dense)
    d2.Product(d,d)
    q := new(mat.Dense)
    q.Kronecker(GenerateI(n), d2)
    expect := new(mat.Dense)
    expect.Product(q,x)
    if !mat.Equal(expect,res){
        t.Fatal("calQx wrong")
    }
    if debug{
        utils.PrintMatrix(expect.T())
        utils.PrintMatrix(res.T())
    }

}

func TestBtCBx(t *testing.T){
    m := 9
    n := 4
    lvl := 1
    N := utils.CalN(lvl)

    avl := make([]float64, m*n)
    for i := range avl{
        avl[i] = randv()
    }
    tvl := make([]float64, m * m)
    for i:=0; i < m; i++{
        tvl[i*m+i] = randv()
    }
    yvl := make([]float64, m*N)
    for i := range yvl{
        yvl[i] = randv()
    }
    ma := mat.NewDense(m,n,avl)
    mt := mat.NewDense(m,m,tvl)

    xvl := make([]float64, n*N)
    for i := range xvl{
        xvl[i] = randv()
    }
    sl := NewIteSolu(ma,mt,m,n,N,yvl,3)
    x := mat.NewVecDense(n*N,xvl)
    res := sl.calBtCBx(x)

    IMat := GenerateI(N)
    B := new(mat.Dense)
    B.Kronecker(ma,IMat)
    C := new(mat.Dense)
    C.Kronecker(mt,IMat)
    expect := new(mat.Dense)
    expect.Product(B.T(),C,B, x)
    if !mat.Equal(expect,res){
        t.Fatal("calBtCBx wrong")
    }
    if debug{
        utils.PrintMatrix(expect.T())
        utils.PrintMatrix(res.T())
    }

}
